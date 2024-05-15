package ratelimit

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"snakeweb/pkg/logger"
)

func Test_slidingWindowLimiter_Allow(t *testing.T) {

	logger.NewLogger("./test")
	runtime.GOMAXPROCS(runtime.NumCPU())

	const collection = "rewards_sliding"
	cli := getLocalRedisCli()

	limiter, _ := NewSlidingWindowLimiter(cli, time.Minute, 6)
	for i := 0; i < 6; i++ {
		allow, err := limiter.Allow(collection, fmt.Sprintf("req_%v", i))
		assert.Nil(t, err)
		assert.True(t, allow)
		time.Sleep(time.Second * 5)
	}

	allow, err := limiter.Allow(collection, fmt.Sprintf("req_%v", 6))
	assert.Nil(t, err)
	assert.False(t, allow)
	time.Sleep(time.Second * 31)

	allow, err = limiter.Allow(collection, fmt.Sprintf("req_%v", 7))
	assert.Nil(t, err)
	assert.True(t, allow)

	card := cli.ZCard(collection).Val()
	assert.Equal(t, int64(6), card)
}
