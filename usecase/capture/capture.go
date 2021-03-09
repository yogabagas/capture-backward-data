package usecase

import (
	"context"
	rds "my-github/capture-backward-data/datastore/redis"
	"my-github/capture-backward-data/domain/model"
	"my-github/capture-backward-data/domain/repository"
	sv "my-github/capture-backward-data/domain/service"
	"time"
)

type CaptureDataImpl struct {
	postgre repository.PostgreRepository
	mongo   repository.MongoRepository
	redis   rds.InternalRedis
}

type Option func(impl *CaptureDataImpl)

func CacheCaptureData(cache rds.InternalRedis) Option {
	return func(impl *CaptureDataImpl) {
		impl.redis = cache
	}
}

func MongoCaptureData(mongo repository.MongoRepository) Option {
	return func(impl *CaptureDataImpl) {
		impl.mongo = mongo
	}
}

func PostgreCaptureData(postgre repository.PostgreRepository) Option {
	return func(impl *CaptureDataImpl) {
		impl.postgre = postgre
	}
}

func NewCaptureConnection(ops ...Option) sv.CaptureData {

	c := &CaptureDataImpl{}

	for _, opt := range ops {
		opt(c)
	}
	return c
}

func (c *CaptureDataImpl) ReadDataAWB(ctx context.Context, from, to time.Time) ([]model.AWBDetailPartner, error) {
	return nil, nil
}

func (c *CaptureDataImpl) InsertDataAWB(ctx context.Context, awb []model.AWBDetailPartner) error {
	return nil
}
