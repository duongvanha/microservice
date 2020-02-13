package handler

import (
	"context"
	micro_models "microservice/pkg/models"

	"github.com/micro/go-micro/v2/util/log"
)

type BuildMovie struct{}

func (e *BuildMovie) Create(ctx context.Context, req *micro_models.Movie, rsp *micro_models.Movie) error {
	log.Log("Received BuildMovie.Call request")

	rsp.Id = 12312
	return nil
}
