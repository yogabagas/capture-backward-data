package redis

import (
	"context"
	"time"

	"github.com/go-redis/cache/v7"
	// "github.com/go-redis/redis"
	"github.com/go-redis/redis/v7"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v7"
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

var ctx = context.Background()

func NewRedisClient(host map[string]string, pass string, db int) InternalRedis {

	client := redis.NewRing(&redis.RingOptions{
		Addrs:        host,
		Password:     pass,
		DB:           db,
		DialTimeout:  time.Duration(30) * time.Second,
		WriteTimeout: time.Duration(30) * time.Second,
		ReadTimeout:  time.Duration(30) * time.Second,
	})

	client.AddHook(nrredis.NewHook(nil))

	if _, err := client.WithContext(ctx).Ping().Result(); err != nil {
		panic(err)
	}

	return &InternalRedisImpl{
		rdb: client,
		cache: &cache.Codec{
			Redis: client,
			Marshal: func(v interface{}) ([]byte, error) {
				return msgpack.Marshal(v)
			},
			Unmarshal: func(b []byte, v interface{}) error {
				return msgpack.Unmarshal(b, v)
			},
		},
	}
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

// func NewInternalRedisImpl(rdb *redis.Ring, expirationTime int64) InternalRedis {
// 	var durationTTL time.Duration
// 	durationTTL = time.Duration(expirationTime)
// 	return &InternalRedisImpl{
// 		rdb: rdb,
// 		cache: &cache.Codec{
// 			Redis: rdb,
// 			Marshal: func(v interface{}) ([]byte, error) {
// 				return msgpack.Marshal(v)
// 			},
// 			Unmarshal: func(b []byte, v interface{}) error {
// 				return msgpack.Unmarshal(b, v)
// 			},
// 		},
// 		ttl: time.Hour * durationTTL,
// 	}
// }
