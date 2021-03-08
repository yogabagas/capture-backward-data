package usecase

import (
	"context"
)

type CaptureData interface {
	Capture(ctx context.Context, from, to string) error
}

type CaptureConn struct {
	mongo string
	redis string
}

func NewCaptureConnection(mongo, redis string) CaptureData {
	return &CaptureConn{mongo: mongo, redis: redis}
}
