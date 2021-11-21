package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Environment string      `json:"env"`
	Redis       RedisConfig `json:"redis"`
}

type RedisConfig struct {
	Host      string `json:"host"`
	Password  string `json:"password"`
	ScanCount int64  `json:"scan_count"`
}

func initConfig(ctx context.Context, configName string) (*Config, error) {
	jsonFile, err := os.Open(configName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	c := new(Config)

	if err := json.Unmarshal(byteValue, &c); err != nil {
		return nil, err
	}

	return c, nil
}
