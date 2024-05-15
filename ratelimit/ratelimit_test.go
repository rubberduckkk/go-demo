package ratelimit

import (
	"testing"

	"github.com/go-redis/redis"
)

func TestCleanUp(t *testing.T) {
	getLocalRedisCli().Del("rewards_sliding")
	getLocalRedisCli().Del("rewards_fixed")
}

func getLocalRedisCli() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
}
