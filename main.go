package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Must provide a path to a config file.")
		return
	}
	fileFromArgs := os.Args[1]

	cfg, err := readConfig(fileFromArgs)
	if err != nil {
		fmt.Println(err)
		return
	}
	splitClient := initSdk(*cfg)

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
