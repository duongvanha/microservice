package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	buildMovie "buildMovie/proto/buildMovie"
)

type BuildMovie struct{}

func (e *BuildMovie) Handle(ctx context.Context, msg *buildMovie.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *buildMovie.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
