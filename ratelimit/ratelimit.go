package ratelimit

import (
	"errors"

	"github.com/go-redis/redis"
)

type Limiter interface {
	Allow(resourceKey, requestKey string) (bool, error)
}

const (
	reject = iota
	allow
)

var (
	ErrNilRedisClient             = errors.New("redis client is nil")
	ErrUnsupportedFixedWindowSize = errors.New("unsupported fixed window size")
)

// checkScript
//
// key: zset 的 key
//
// cutoff: unix 毫秒级时间戳，脚本会清理 cutoff 之前的数据
//
// allowCnt: 允许存在的数据个数
//
// callKey: 单次请求数据的唯一 ID
//
// now: unix 毫秒级时间戳
var checkScript = redis.NewScript(`
local key = KEYS[1]
local cutoff = ARGV[1]
local allowCnt = ARGV[2]
local callKey = ARGV[3]
local now = ARGV[4]

redis.call('ZREMRANGEBYSCORE', key, 0, cutoff)

local reqCnt = redis.call('ZCARD', key)

if reqCnt >= tonumber(allowCnt) then
	return 0
else
	redis.call('ZADD', key, now, callKey)
	return 1
end
`)
