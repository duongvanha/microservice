package handler

import (
	"context"
	micro_models "microservice/src/pkg/models"

	"github.com/micro/go-micro/util/log"
)

type BuildMovie struct{}

func (e *BuildMovie) Create(ctx context.Context, req *micro_models.Movie, rsp *micro_models.Movie) error {
	log.Log("Received BuildMovie.Call request")

	return nil
}
