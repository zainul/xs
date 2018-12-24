package redis

import (
	"github.com/zainul/xs/internal/pkg/cache"
)

type cacheImpl struct {
}

func NewCached(config string) cache.Cache {
	return &cacheImpl{}
}

// Get key is the key , and field is optional value
func (c *cacheImpl) Get(key string, field ...string) ([]byte, error) {
	return nil, nil
}

// Set key is the key , field (optional) and value must be added
func (c *cacheImpl) Set(key string, fieldAndValue ...string) (bool, error) {
	return false, nil
}

// Delete is deleted the key, and field is optional
func (c *cacheImpl) Del(key string, field ...string) (int64, error) {
	return 0, nil
}

// Set expired for the cache
func (c *cacheImpl) Setex(key string, seconds int, value string) error {
	return nil
}

// Exists key is the key and field is optional
func (c *cacheImpl) Exists(key string, field ...string) (bool, error) {
	return false, nil
}

// Expire key is the key and second is parameter to be expired in unit (second)
func (c *cacheImpl) Expire(key string, seconds int) (bool, error) {
	return false, nil
}

// Expire key is the key and timestamp is parameter to be expired in unit (timestamp)
func (c *cacheImpl) ExpireAt(key string, timestamp int64) (bool, error) {
	return false, nil
}

// TTL
func (c *cacheImpl) TTL(key string) (int64, error) {
	return 0, nil
}
