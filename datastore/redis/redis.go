package redis

import (
	"context"
	"time"

	"github.com/go-redis/cache/v7"
	"github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack"
)

type InternalRedisImpl struct {
	rdb   *redis.Ring
	cache *cache.Codec
	ttl   time.Duration
}

type InternalRedis interface {
	Get(ctx context.Context, key string, obj interface{}) error
	Create(ctx context.Context, key string, expirationTimeInHour int64, value interface{}) error
}

func (ar *InternalRedisImpl) Get(ctx context.Context, key string, obj interface{}) error {

	if err := ar.cache.GetContext(ctx, key, obj); err != nil {
		return err
	}
	return nil
}

func (ar *InternalRedisImpl) Create(ctx context.Context, key string, expirationTime int64, value interface{}) error {
	if expirationTime > 0 {
		ar.ttl = time.Duration(expirationTime) * time.Hour
	}
	if err := ar.cache.Set(&cache.Item{
		Ctx:        ctx,
		Key:        key,
		Object:     value,
		Expiration: ar.ttl,
	}); err != nil {
		return err
	}
	return nil
}

func NewInternalRedisImpl(rdb *redis.Ring, expirationTime int64) InternalRedis {
	var durationTTL time.Duration
	durationTTL = time.Duration(expirationTime)
	return &InternalRedisImpl{
		rdb: rdb,
		cache: &cache.Codec{
			Redis: rdb,
			Marshal: func(v interface{}) ([]byte, error) {
				return msgpack.Marshal(v)
			},
			Unmarshal: func(b []byte, v interface{}) error {
				return msgpack.Unmarshal(b, v)
			},
		},
		ttl: time.Hour * durationTTL,
	}
}
