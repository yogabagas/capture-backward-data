package command

import (
	"log"
	"my-github/capture-backward-data/controller"
	usecase "my-github/capture-backward-data/usecase/capture"
	"net/http"
)

var Run = func() {
	cc := NewCaptureController()

	mux := http.NewServeMux()
	mux.HandleFunc("/", cc.CaptureAWB)
	if err := http.ListenAndServe(":8111", mux); err != nil {
		log.Fatal(err)
	}
}

func NewCaptureController() controller.CaptureDataController {
	return controller.NewCaptureData(usecase.NewCaptureConnection(Options()...))
}

func Options() []usecase.Option {
	var Opts []usecase.Option
	Opts = append(Opts,
		usecase.CacheCaptureData(InitRedis()),
		usecase.PostgreCaptureData(InitPostgre()),
		usecase.MongoCaptureData(InitMongo()))
	return Opts
}
