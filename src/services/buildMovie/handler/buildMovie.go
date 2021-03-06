package handler

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/logger"
	micro_models "microservice/src/gopkg/models"
	"time"
)

type BuildMovie struct {
	Broker broker.Broker
}

var i = 0

func (e *BuildMovie) Create(ctx context.Context, req *micro_models.Movie, rsp *micro_models.Movie) error {
	logger.Infof("Received BuildMovie.Call request")

	err := e.Broker.Publish("capture_additional_charge", &broker.Message{
		Header: map[string]string{
			"id": fmt.Sprintf("%d", i),
		},
		Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
	})

	i++

	if err != nil {
		logger.Error("public error: %v", err)
	}

	rsp.CreatedAt = time.Now().Unix()
	return nil
}
