package main

import (
	"context"
	"errors"
	"fmt"

	redis "github.com/go-redis/redis/v8"
	"github.com/tokopedia/tdk/go/log"
)

func connectRedis(ctx context.Context, redisHost, redisPassword string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
	})

	result, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	if result != "PONG" {
		return nil, errors.New("result is not \"PONG\"")
	}

	return rdb, nil
}

func scanRedis(ctx context.Context, rdb *redis.Client, f *File, keyword string, scanCount int64) error {
	result := make([]string, 0)
	cursor := uint64(0)

	log.Infof("Looking up for \"%s\"", keyword)

	for {
		keys, curr, err := rdb.Scan(ctx, cursor, keyword, scanCount).Result()
		if err != nil {
			return err
		}
		cursor = curr
		result = append(result, keys...)

		for _, key := range keys {
			if err := f.Write(fmt.Sprintf("%s\n", key)); err != nil {
				log.Error("Failed to write file:", err)
			}
		}
		if err := f.Sync(); err != nil {
			log.Error("Failed to sync file:", err)
		}

		log.Info("Current data: ", "result count:", len(result), " cursor:", cursor)

		if cursor == 0 {
			break
		}
	}

	return nil
}
