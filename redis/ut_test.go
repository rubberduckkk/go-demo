package main

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
)

func buildMiniRedis(t *testing.T) (*miniredis.Miniredis, *redis.Client) {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}

	cli := redis.NewClient(&redis.Options{
		Addr:        s.Addr(),
		Password:    "",
		PoolSize:    5,
		DB:          0,
		IdleTimeout: time.Hour,
		MaxRetries:  2,
	})
	return s, cli
}
