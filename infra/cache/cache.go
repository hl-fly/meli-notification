package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hector-leite/meli-notification/domain/constants"
	"github.com/hector-leite/meli-notification/domain/contract"
)

// RedisCache implements the CacheManager interface
type RedisCache struct {
	redis             *redis.Client
	prefix            string
	defaultExpiration time.Duration
}

// Connect returns a new connection, representing a CacheManager.
func Connect(
	host string,
	port int,
	db int,
	pass string,
	prefix string,
	defaultExpiration time.Duration,
) contract.CacheManager {

	return &RedisCache{
		prefix:            prefix,
		defaultExpiration: defaultExpiration,
		redis: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Password: pass,
			DB:       db,
		}),
	}
}

func (r *RedisCache) buildKey(key string) string {
	if r.prefix != "" {
		return r.prefix + ":" + key
	}

	return key
}

func (r *RedisCache) getItem(key string) (data []byte, err error) {
	val, err := r.redis.Get(context.Background(), r.buildKey(key)).Bytes()
	if err == redis.Nil {
		return val, constants.ErrCacheMiss
	}
	if err != nil {
		return val, err
	}

	return val, nil
}

func (r *RedisCache) setItem(key string, data []byte) error {
	err := r.redis.Set(context.Background(), r.buildKey(key), data, r.defaultExpiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCache) GetStruct(key string, data interface{}) (err error) {
	val, err := r.getItem(key)
	if err == constants.ErrCacheMiss {
		return constants.ErrCacheMiss
	}
	if err != nil {
		return err
	}

	err = json.Unmarshal(val, &data)
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCache) SetStruct(key string, data interface{}) error {
	dataString, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = r.setItem(key, dataString)
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCache) Invalidate(key string) error {
	err := r.redis.Del(context.Background(), r.buildKey(key)).Err()
	if err == redis.Nil {
		return constants.ErrCacheMiss
	}
	if err != nil {
		return err
	}

	return nil
}
