package handler

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"
	micro_models "microservice/src/gopkg/models"
	"time"
)

type BuildMovie struct{}

func (e *BuildMovie) Create(ctx context.Context, req *micro_models.Movie, rsp *micro_models.Movie) error {
	log.Log("Received BuildMovie.Call request")

	rsp.CreatedAt = time.Now().Unix()
	return nil
}
