package service

import (
	"context"
	"time"
)

type CaptureData interface {
	ReadDataAWB(ctx context.Context, from, to time.Time) error

	InsertDataAWB(ctx context.Context) error
}
