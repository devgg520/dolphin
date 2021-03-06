package persistence

import (
	"strconv"
	"time"

	"github.com/2637309949/dolphin/packages/cache/utils"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/redis"
)

// RedisStore represents the cache with redis persistence
type RedisStore struct {
	redisClient       *redis.Client
	defaultExpiration time.Duration
}

// NewRedisCache returns a RedisStore
// until redigo supports sharding/clustering, only one host will be in hostList
func NewRedisCache(addr string, password string, db int, defaultExpiration time.Duration) *RedisStore {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if _, err := redisClient.Ping().Result(); err != nil {
		logrus.Error("Redis: connect failed")
	}
	return &RedisStore{redisClient, defaultExpiration}
}

// NewRedisCacheWithPool returns a RedisStore using the provided pool
// until redigo supports sharding/clustering, only one host will be in hostList
func NewRedisCacheWithPool(redisClient *redis.Client, defaultExpiration time.Duration) *RedisStore {
	return &RedisStore{redisClient, defaultExpiration}
}

// Set (see CacheStore interface)
func (c *RedisStore) Set(key string, value interface{}, expires time.Duration) error {
	return c.invoke(key, value, expires)
}

// Add (see CacheStore interface)
func (c *RedisStore) Add(key string, value interface{}, expires time.Duration) error {
	return c.invoke(key, value, expires)
}

// Replace (see CacheStore interface)
func (c *RedisStore) Replace(key string, value interface{}, expires time.Duration) error {
	err := c.invoke(key, value, expires)
	if value == nil {
		return ErrNotStored
	}
	return err
}

// Get (see CacheStore interface)
func (c *RedisStore) Get(key string, ptrValue interface{}) error {
	buf, err := c.redisClient.Get(key).Bytes()
	if err != nil {
		return err
	}
	return utils.Deserialize(buf, ptrValue)
}

func exists(client *redis.Client, key string) bool {
	iresult := client.Exists(key)
	if err := iresult.Err(); err != nil && err != redis.Nil {
		return false
	}
	return iresult.Val() == 0
}

// Delete (see CacheStore interface)
func (c *RedisStore) Delete(key string) error {
	if !exists(c.redisClient, key) {
		return ErrCacheMiss
	}
	result := c.redisClient.Del(key)
	if err := result.Err(); err != nil && err != redis.Nil {
		return err
	}
	return nil
}

// Increment (see CacheStore interface)
func (c *RedisStore) Increment(key string, delta uint64) (uint64, error) {
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that. Since we need to do increment
	// ourselves instead of natively via INCRBY (redis doesn't support wrapping), we get the value
	// and do the exists check this way to minimize calls to Redis
	val, err := c.redisClient.Get(key).Result()
	if val == "" {
		return 0, ErrCacheMiss
	}
	if err == nil {
		currentVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, err
		}
		sum := currentVal + int64(delta)
		_, err = c.redisClient.Set(key, sum, 0).Result()
		return uint64(sum), nil
	}

	return 0, err
}

// Decrement (see CacheStore interface)
func (c *RedisStore) Decrement(key string, delta uint64) (newValue uint64, err error) {
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that, hence the exists call
	if !exists(c.redisClient, key) {
		return 0, ErrCacheMiss
	}
	// Decrement contract says you can only go to 0
	// so we go fetch the value and if the delta is greater than the amount,
	// 0 out the value
	val, err := c.redisClient.Get(key).Result()
	if val == "" {
		return 0, ErrCacheMiss
	}
	currentVal, err := strconv.ParseInt(val, 10, 64)
	if err == nil && delta > uint64(currentVal) {
		tempint, err := c.redisClient.DecrBy(key, currentVal).Result()
		return uint64(tempint), err
	}
	tempint, err := c.redisClient.DecrBy(key, int64(delta)).Result()
	return uint64(tempint), err
}

// Flush (see CacheStore interface)
func (c *RedisStore) Flush() error {
	_, err := c.redisClient.FlushAll().Result()
	return err
}

func (c *RedisStore) invoke(key string, value interface{}, expires time.Duration) error {

	switch expires {
	case DEFAULT:
		expires = c.defaultExpiration
	case FOREVER:
		expires = time.Duration(0)
	}

	b, err := utils.Serialize(value)
	if err != nil {
		return err
	}

	if expires > 0 {
		_, err := c.redisClient.SetNX(key, b, expires).Result()
		return err
	}

	_, err = c.redisClient.Set(key, b, expires).Result()
	return err

}
