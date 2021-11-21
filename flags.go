package main

import (
	"context"
	"errors"
	"flag"
)

type Flag struct {
	ConfigFile string
	Keyword    string
}

func initFlags(ctx context.Context) (*Flag, error) {
	configFile := flag.String("config", "config.json", "config file name")
	keyword := flag.String("keyword", "", "keyword to search")

	flag.Parse()

	if len(*keyword) == 0 {
		return nil, errors.New("keyword param is not valid")
	}

	return &Flag{ConfigFile: *configFile, Keyword: *keyword}, nil
}
