package service

import "context"

type CaptureData interface {
	CaptureDataBackward(ctx context.Context, from, to string) error
}
