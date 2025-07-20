package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type Config struct {
	ApiKey    string       `json:"apiKey"`
	IsStaging bool         `json:"isStaging"`
	Flags     []FlagConfig `json:"flags"`
}

type FlagConfig struct {
	Name        string         `json:"name"`
	Impressions int            `json:"impressions"`
	Events      []EventsConfig `json:"events"`
}

type EventsConfig struct {
	EventType   string                        `json:"eventType"`
	TrafficType string                        `json:"trafficType"`
	Treatments  map[string]EventValueSettings `json:"treatments"`
}

type EventValueSettings struct {
	Value      *int                   `json:"value"`
	Count      *int                   `json:"count"`
	Properties map[string]interface{} `json:"properties"`
}

func readConfig(path string) (*Config, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to read config file %s", path))
	}
	defer jsonFile.Close()

	var cfg Config
	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &cfg)

	fmt.Println(cfg)

	return &cfg, nil
}
