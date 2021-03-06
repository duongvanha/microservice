package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	micro_models "microservice/src/gopkg/models"
	MovieRepository "microservice/src/gopkg/services"
)

func main() {

	service := micro.NewService(micro.Name("shippy.service.consignment"))

	client := service.Client()

	movieRepository := MovieRepository.NewMovieRepositoryService("shippy.service.consignment", client)

	movie, err := movieRepository.Create(context.TODO(), &micro_models.Movie{Id: 1})

	println(movie, err)
}
