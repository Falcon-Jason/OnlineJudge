package dao

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var Redis *redis.Client

func connectRedis(uri string) error {
	if Redis != nil {
		return errors.New("redis has been connected")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	opt, err := redis.ParseURL(uri)
	if err != nil {
		return err
	}
	rdb := redis.NewClient(opt)
	if err = rdb.Ping(ctx).Err(); err != nil {
		return err
	}

	Redis = rdb
	return nil
}

func closeRedis() {
	if Redis == nil {
		return
	}

	err := Redis.Close()
	if err != nil {
		log.Fatalf("failed to close redis: %v", err)
	}
}
