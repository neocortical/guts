package omnicache

import (
	"time"

	"github.com/Hearst-DD/cache"
)

type omnicache struct {
	readCache   cache.Cache
	writeCaches []cache.Cache
}

// New returns a new Redis-backed cache.
func New(readCache cache.Cache, writeCaches []cache.Cache) cache.Cache {
	return &omnicache{
		readCache:   readCache,
		writeCaches: writeCaches,
	}
}

func (oc *omnicache) Put(key string, value interface{}, ttl time.Duration) {
	for _, c := range oc.writeCaches {
		c.Put(key, value, ttl)
	}
}

func (oc *omnicache) Get(key string) (value interface{}, status cache.Result) {
	if oc.readCache != nil {
		return oc.readCache.Get(key)
	}

	return value, cache.NotFound
}

func (oc *omnicache) Size() int {
	return -1
}
