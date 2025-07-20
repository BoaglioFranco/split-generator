package main

import (
	"fmt"
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

