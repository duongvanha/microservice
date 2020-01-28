package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	movieRepository "microservice/pkg/services"
	"microservice/services/buildMovie/handler"
	"microservice/services/buildMovie/subscriber"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	_ = movieRepository.RegisterMovieRepositoryHandler(service.Server(), new(handler.BuildMovie))

	// Register Struct as Subscriber
	_ = micro.RegisterSubscriber("shippy.service.consignment", service.Server(), new(subscriber.BuildMovie))

	// Register Function as Subscriber
	_ = micro.RegisterSubscriber("shippy.service.consignment", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
