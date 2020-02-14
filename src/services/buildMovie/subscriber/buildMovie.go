package subscriber

import (
	"context"
	micro_models "microservice/src/pkg/models"
)

type BuildMovie struct{}

func (e *BuildMovie) Handle(ctx context.Context, msg *micro_models.Movie) error {
	println("Handler Received message: ", msg.Id)
	return nil
}
func (e *BuildMovie) Process(ctx context.Context, msg *micro_models.Movie) error {
	println("Handler Received message: ", msg.Id)
	return nil
}

func Handler(ctx context.Context, msg *micro_models.Movie) error {
	println("Function Received message: ", msg.Id)
	return nil
}
