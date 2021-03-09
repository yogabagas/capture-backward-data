package repository

import (
	"context"
	"my-github/capture-backward-data/domain/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type PostgreRepository interface {
	ReadDataInterval(ctx context.Context, from, to string) ([]model.AWBDetailPartner, error)
}

type MongoRepository interface {
	AddBulkInsert(ctx context.Context, collectionName string, operation mongo.WriteModel) error
	RunBulkInsert(ctx context.Context) error
}
