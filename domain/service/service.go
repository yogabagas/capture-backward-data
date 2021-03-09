package service

import (
	"context"
	"my-github/capture-backward-data/domain/model"
	"time"
)

type CaptureData interface {
	ReadDataAWB(ctx context.Context, from, to time.Time) ([]model.AWBDetailPartner, error)

	InsertDataAWB(ctx context.Context, awb []model.AWBDetailPartner) error
}
