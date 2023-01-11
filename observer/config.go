package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Observer struct {
		Freq            int `json:"freq"`
		StartNotifyDate int `json:"startNotifyDate"`
		EndNotifyDate   int `json:"endNotifyDate"`
	} `json:"observer"`
	ValuableHours []string `json:"valuableHours"`
	Subscribers   []string `json:"subscribers"`
}

func loadConfig() Config {
	var cfg Config
	bytes, _ := os.ReadFile("config.json")
	json.Unmarshal(bytes, &cfg)

	return cfg
}