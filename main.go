package main

import (
	"fmt"
	"strconv"
	"math/rand"
    "github.com/splitio/go-client/v6/splitio/client"
    "github.com/splitio/go-client/v6/splitio/conf"
)


func initSdk(cfg Config) *client.SplitClient {

	sdkSettings := conf.Default()

	if cfg.IsStaging {
		sdkSettings.Advanced.SdkURL = "https://sdk.split-stage.io/api"
		sdkSettings.Advanced.EventsURL = "https://events.split-stage.io/api"
	}

	// sdkSettings.LoggerConfig.LogLevel = logging.LevelAll
	factory, err := client.NewSplitFactory(cfg.ApiKey, sdkSettings)
	if err != nil {
		fmt.Printf("Error connecting to the SDK")
	}
	client := factory.Client()
	err = client.BlockUntilReady(5)
	if err != nil {
		fmt.Printf("SDK init error: %s\n", err)
	}
	return client
}

func main() {
	cfg := readConfig()

	splitClient := initSdk(cfg)

	for _, flag := range cfg.Flags {
		randSeed := rand.Int()

		for j := range flag.Impressions {
			key := strconv.Itoa(j) + "-" + strconv.Itoa(randSeed)

			treatment := splitClient.Treatment(key, flag.Name, nil)
			fmt.Println(flag.Name, "  ", j, "    ", treatment)

			for _, e := range flag.Events {
				eventCfg := e.Treatments[treatment]

				var value interface{}
				if eventCfg.Value != nil {
					value = *eventCfg.Value
				}
				if eventCfg.Count == nil {
					splitClient.Track(key, e.TrafficType, e.EventType, value, eventCfg.Properties)
					fmt.Printf("Sent event %v (%v) for key %v -- Value: %v \n", e.EventType, e.TrafficType, key, value)
				} else {
					for range *eventCfg.Count {

						splitClient.Track(key, e.TrafficType, e.EventType, value, eventCfg.Properties)
						fmt.Printf("Sent event %v (%v) for key %v -- Value: %v \n", e.EventType, e.TrafficType, key, eventCfg.Value)

					}
				}

			}

		}
	}

	splitClient.Destroy()
}
