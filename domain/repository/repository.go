package repository

import (
	"context"
	"my-github/capture-backward-data/domain/model"
)

type PostgreRepository interface {
	ReadDataInterval(ctx context.Context, from, to string) ([]model.AWBDetailPartner, error)
}

type MongoRepository interface {
	BulkInsertDataAWB(ctx context.Context, awb []model.AWBDetailPartner) error
}
