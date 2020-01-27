package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"microservice/src/services/buildMovie/handler"
	"microservice/src/services/buildMovie/subscriber"

	movieRepository "microservice/src/pkg/services"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("ha.duong.srv.buildMovie"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	_ = movieRepository.RegisterMovieRepositoryHandler(service.Server(), new(handler.BuildMovie))

	// Register Struct as Subscriber
	_ = micro.RegisterSubscriber("ha.duong.srv.buildMovie", service.Server(), new(subscriber.BuildMovie))

	// Register Function as Subscriber
	_ = micro.RegisterSubscriber("ha.duong.srv.buildMovie", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
