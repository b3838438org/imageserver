// Redis cache
package redis

import (
	redigo "github.com/garyburd/redigo/redis"
	"github.com/pierrre/imageserver"
	"strconv"
	"time"
)

// Uses Gary Burd's Redis client https://github.com/garyburd/redigo
type RedisCache struct {
	Pool *redigo.Pool

	Expire time.Duration // optional
}

func (redis *RedisCache) Get(key string, parameters imageserver.Parameters) (*imageserver.Image, error) {
	data, err := redis.getData(key, parameters)
	if err != nil {
		return nil, imageserver.NewCacheMissError(key, redis, err)
	}

	image, err := imageserver.NewImageUnmarshalBinary(data)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func (redis *RedisCache) getData(key string, parameters imageserver.Parameters) ([]byte, error) {
	conn := redis.Pool.Get()
	defer conn.Close()

	return redigo.Bytes(conn.Do("GET", key))
}

func (redis *RedisCache) Set(key string, image *imageserver.Image, parameters imageserver.Parameters) error {
	data, err := image.MarshalBinary()
	if err != nil {
		return err
	}

	params := []interface{}{key, data}

	if redis.Expire != 0 {
		params = append(params, "EX", strconv.Itoa(int(redis.Expire.Seconds())))
	}

	conn := redis.Pool.Get()
	defer conn.Close()

	_, err = conn.Do("SET", params...)
	if err != nil {
		return err
	}

	return nil
}
