package controller

import (
	"context"
	"my-github/capture-backward-data/domain/service"
	"net/http"
	"time"
)

type CaptureData struct {
	capture service.CaptureData
}

type CaptureDataController interface {
	CaptureAWB(w http.ResponseWriter, r *http.Request)
}

func NewCaptureData(capture service.CaptureData) CaptureDataController {
	return &CaptureData{capture: capture}
}

func (c *CaptureData) CaptureAWB(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	awb, err := c.capture.ReadDataAWB(context.Background(), now, now.AddDate(0, -3, 0))
	if err != nil {
		return
	}

	c.capture.InsertDataAWB(context.Background(), awb)
}
