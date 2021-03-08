package usecase

import (
	"context"
	sv "my-github/capture-backward-data/domain/service"
	rds "my-github/capture-bakward-data/datastore/redis"
	repo "my-github/capture-bakward-data/domain/repository"
)

type CaptureDataImpl struct {
	sql   repo.PostgreRepository
	mongo repo.MongoRepository
	redis rds.InternalRedis
}

func NewCaptureConnection(sql, mongo, redis string) sv.CaptureData {
	return &CaptureDataImpl{sql: sql, mongo: mongo, redis: redis}
}

func (c *CaptureDataImpl) CaptureDataBackward(ctx context.Context, from, to string) error {
	return nil
}
