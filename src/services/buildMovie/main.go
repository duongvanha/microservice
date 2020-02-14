package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	micro_models "microservice/src/pkg/models"
	micro_services "microservice/src/pkg/services"
	"microservice/srcservices/buildMovie/handler"
	"microservice/srcservices/buildMovie/subscriber"
	"time"
)

// send events using the publisher
func sendEv(topic string, p micro.Publisher) {
	t := time.NewTicker(time.Second)

	for _ = range t.C {
		// create new event
		ev := &micro_models.Movie{
			Id: time.Now().Unix(),
		}
		if err := p.Publish(context.Background(), ev); err != nil {
			log.Logf("error publishing: %v", err)
		}
	}
}

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	_ = micro_services.RegisterMovieRepositoryHandler(service.Server(), new(handler.BuildMovie))

	go func() {
		// Register Struct as Subscriber
		e2 := micro.RegisterSubscriber("shippy.cli.consignment", service.Server(), new(subscriber.BuildMovie))

		// Register Function as Subscriber
		e3 := micro.RegisterSubscriber("shippy.cli.consignment", service.Server(), subscriber.Handler)

		println(e2, e3)
	}()

	pub1 := micro.NewEvent("shippy.cli.consignment", service.Client())

	// pub to topic 1
	go sendEv("shippy.cli.consignment", pub1)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
