package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tokopedia/tdk/go/log"
)

func main() {
	var (
		ctx = context.Background()
	)

	initLog(ctx)
	log.Info("Logger initialized.")

	flags, err := initFlags(ctx)
	if err != nil {
		log.Fatal("Flag is invalid:", err)
	}

	cfg, err := initConfig(ctx, flags.ConfigFile)
	if err != nil {
		log.Fatal("Failed to init config:", err)
	}
	log.Info("Config initialized. Env:", cfg.Environment)

	f, err := CreateFile(
		fmt.Sprintf("result-%s.txt", time.Now().Format("2006_01_02-15_04_05")),
	)
	defer f.Defer()
	if err != nil {
		log.Fatal("Failed to create file:", err)
	}

	rdb, err := connectRedis(ctx, cfg.Redis.Host, cfg.Redis.Password)
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	if err := scanRedis(ctx, rdb, f, flags.Keyword, cfg.Redis.ScanCount); err != nil {
		log.Fatal("Scan returns error", err)
	}

	log.Infof("Scan finished. Result is saved in %s.", f.fileName)
}
