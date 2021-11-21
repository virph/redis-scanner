package main

import (
	"context"

	"github.com/tokopedia/tdk/go/log"
)

func initLog(ctx context.Context) {
	err := log.SetConfig(&log.Config{
		Level:   "debug",
		AppName: "redis-scanner",
		Caller:  true,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.InitLogContext(ctx)
}
