package goldilock

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// Locker interface defining the functionality of acquiring a lock.
type Locker interface {
	Lock(key string) bool
}

type locker struct {
	redisPool *redis.Pool
	ttl       int64
}

// NewLocker creates a new concrete Locker. Note that TTL will be truncted to the nearest millisecond.
func NewLocker(pool *redis.Pool, ttl time.Duration) Locker {
	return &locker{
		redisPool: pool,
		ttl:       ttl.Nanoseconds() / 1000000,
	}
}

func (l *locker) Lock(key string) bool {
	var redisConn = l.redisPool.Get()
	defer redisConn.Close()
	response, err := redis.String(redisConn.Do("SET", key, 1, "NX", "PX", l.ttl))
	return err == nil && response == "OK"
}
