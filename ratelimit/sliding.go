package ratelimit

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type slidingWindowLimiter struct {
	cli        *redis.Client
	windowSize time.Duration
	allowCnt   int
}

// NewSlidingWindowLimiter
//
// 限流逻辑: windowSize 时间段内最多允许 allowCount 次请求
func NewSlidingWindowLimiter(redisCli *redis.Client, windowSize time.Duration, allowCount int) (Limiter, error) {
	if redisCli == nil {
		return nil, ErrNilRedisClient
	}
	return &slidingWindowLimiter{
		cli:        redisCli,
		windowSize: windowSize,
		allowCnt:   allowCount,
	}, nil
}

func (l *slidingWindowLimiter) Allow(resourceKey, requestKey string) (bool, error) {
	now := time.Now()
	res, err := checkScript.Run(l.cli, []string{resourceKey},
		now.Add(-l.windowSize).UnixMilli(), l.allowCnt, requestKey, now.UnixMilli()).Int()
	if err != nil {
		logrus.WithError(err).WithField("resource", resourceKey).WithField("request", requestKey).
			Warn("[Limiter] call redis failed")
		return false, err
	}

	return res == allow, nil
}
