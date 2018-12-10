package cache

// Cache is represented of the general cache, should be support various cache in general like :
// - redis
// - memcached
// - aerospike
type Cache interface {
	// Get key is the key , and field is optional value
	Get(key string, field ...string) ([]byte, error)
	// Set key is the key , field (optional) and value must be added
	Set(key string, fieldAndValue ...string) (bool, error)
	// Delete is deleted the key, and field is optional
	Del(key string, field ...string) (int64, error)
	// Set expired for the cache
	Setex(key string, seconds int, value string) error
	// Exists key is the key and field is optional
	Exists(key string, field ...string) (bool, error)
	// Expire key is the key and second is parameter to be expired in unit (second)
	Expire(key string, seconds int) (bool, error)
	// Expire key is the key and timestamp is parameter to be expired in unit (timestamp)
	ExpireAt(key string, timestamp int64) (bool, error)
	// TTL
	TTL(key string) (int64, error)
}
