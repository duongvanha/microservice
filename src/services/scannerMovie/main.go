package main

import (
	"github.com/micro/go-micro/v2/logger"
	micro_app "microservice/src/gopkg/core/microApp"
	"time"
)

func main() {
	// create a service
	micro_app.NewService()

	logger.Info("app running", time.Now().String())

}
