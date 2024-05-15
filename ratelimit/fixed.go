package ratelimit

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type FixedWindowSize int

const (
	WindowSizeMinute FixedWindowSize = iota
	WindowSizeHour
	WindowSizeDay
	WindowSizeMonth
	WindowSizeYear
)

var supportedWindowSize = map[FixedWindowSize]bool{
	WindowSizeMinute: true,
	WindowSizeHour:   true,
	WindowSizeDay:    true,
	WindowSizeMonth:  true,
	WindowSizeYear:   true,
}

func (s FixedWindowSize) String() string {
	var str string
	switch s {
	case WindowSizeMinute:
		str = "minute"
	case WindowSizeHour:
		str = "hour"
	case WindowSizeDay:
		str = "day"
	case WindowSizeMonth:
		str = "month"
	case WindowSizeYear:
		str = "year"
	default:
		str = "unknown"
	}
	return str
}

type fixedWindowRateLimiter struct {
	cli        *redis.Client
	windowSize FixedWindowSize
	allowCnt   int
}

// NewFixedWindowRateLimiter
//
// 限流逻辑: 当前时间的分钟、小时、天、月、年内，最多允许 allowCount 次请求（不包含窗口边界上的请求）
//
// 示例:
//
// 假设当前时间为 2024-05-15 16:30:25, windowSize = WindowSizeHour, allowCount = 5
//
// 限流逻辑: 2024-05-15 16:00:00 至 2024-05-15 16:30:25 最多允许 5 次请求
func NewFixedWindowRateLimiter(redisCli *redis.Client, windowSize FixedWindowSize, allowCount int) (Limiter, error) {
	if redisCli == nil {
		return nil, ErrNilRedisClient
	}

	if _, ok := supportedWindowSize[windowSize]; !ok {
		return nil, ErrUnsupportedFixedWindowSize
	}

	return &fixedWindowRateLimiter{
		cli:        redisCli,
		windowSize: windowSize,
		allowCnt:   allowCount,
	}, nil
}

func (f *fixedWindowRateLimiter) Allow(resourceKey, requestKey string) (bool, error) {
	now := time.Now()
	cutoff := f.calculateCutoff(now)
	res, err := checkScript.Run(f.cli, []string{resourceKey},
		cutoff.UnixMilli(), f.allowCnt, requestKey, now.UnixMilli()).Int()
	if err != nil {
		logrus.WithError(err).WithField("resource", resourceKey).WithField("request", requestKey).
			Warn("[Limiter] call redis failed")
		return false, err
	}

	return res == allow, nil
}

func (f *fixedWindowRateLimiter) calculateCutoff(now time.Time) time.Time {
	var cutoff time.Time
	switch f.windowSize {
	case WindowSizeMinute:
		cutoff = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())
	case WindowSizeHour:
		cutoff = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
	case WindowSizeDay:
		cutoff = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	case WindowSizeMonth:
		cutoff = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	case WindowSizeYear:
		cutoff = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	}
	return cutoff
}
