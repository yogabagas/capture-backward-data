package usecase

import (
	"context"
	rds "my-github/capture-backward-data/datastore/redis"
	"my-github/capture-backward-data/domain/repository"
	sv "my-github/capture-backward-data/domain/service"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (c *CaptureDataImpl) ReadDataAWB(ctx context.Context, from, to time.Time) error {
	var records interface{}

	awb := bson.M{
		"awbNumber":          records,
		"originCode":         records,
		"destinationCode":    records,
		"weight":             records,
		"sender":             records,
		"senderAddress":      records,
		"representativeCode": records,
		"isDeleted":          records,
		"updatedAt":          records,
	}

	status := bson.M{}
	metadata := bson.M{}

	awbForInsert := bson.M{
		"data":     awb,
		"status":   status,
		"metadata": metadata,
	}

	operation := mongo.NewUpdateOneModel()
	operation.SetUpsert(true)
	operation.SetFilter(bson.M{"awb_id": awb})
	operation.SetUpdate(bson.M{
		"$setOnInsert": awbForInsert,
	})

	c.mongo.AddBulkInsert(ctx, "cdcAWB", operation)
	return nil
}

func (c *CaptureDataImpl) InsertDataAWB(ctx context.Context) error {
	return c.mongo.RunBulkInsert(ctx)
}
