package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	micro_models "microservice/pkg/models"
)

type BuildMovie struct{}

func (e *BuildMovie) Handle(ctx context.Context, msg *micro_models.Movie) error {
	log.Log("Handler Received message: ", msg.Id)
	return nil
}

func Handler(ctx context.Context, msg *micro_models.Movie) error {
	log.Log("Function Received message: ", msg.Id)
	return nil
}
