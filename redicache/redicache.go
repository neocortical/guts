package redicache

import (
	"time"

	"github.com/Hearst-DD/cache"
	"github.com/garyburd/redigo/redis"
)

// New returns a new Redis-backed cache.
func New(pool *redis.Pool) cache.Cache {
	return &rediscache{
		redisPool: pool,
	}
}

type rediscache struct {
	redisPool *redis.Pool
}

func (rc *rediscache) Put(key string, value interface{}, ttl time.Duration) {
	var redisConn = rc.redisPool.Get()
	defer redisConn.Close()

	var expireMs = int(ttl.Nanoseconds() / 1000000)

	if expireMs > 0 {
		redisConn.Do("SET", key, value, "PX", expireMs)
	} else {
		redisConn.Do("SET", key, value)
	}
}

func (rc *rediscache) Get(key string) (value interface{}, status cache.Result) {
	var redisConn = rc.redisPool.Get()
	defer redisConn.Close()

	value, err := redisConn.Do("GET", key)
	if err != nil {
		return value, cache.NotFound
	}

	return value, cache.OK
}

func (rc *rediscache) Size() int {
	return -1
}
