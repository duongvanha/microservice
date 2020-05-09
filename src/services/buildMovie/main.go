package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/logger"
	micro_app "microservice/src/gopkg/core/microApp"
	micro_models "microservice/src/gopkg/models"
	micro_services "microservice/src/gopkg/services"
	"microservice/src/services/buildMovie/handler"
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
			logger.Error("error publishing: %v", err)
		}
	}
}

func sub(broker2 broker.Broker) {
	_, err := broker2.Subscribe("capture_additional_charge", func(p broker.Event) error {
		logger.Infof("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	}, broker.Queue("order_capture_additional_charge"))
	if err != nil {
		logger.Error("error", err)
	}
}

func main() {
	// New Service
	service := micro_app.NewService(micro.Name("go.haduong.service.build_movie"))

	service.Client()

	// Initialise service
	service.Init()

	defaultBroker, err := micro_app.GetBroker()
	if err != nil {
		logger.Fatalf("Broker Init error: %v", err)
	}

	_, err = defaultBroker.Subscribe("capture_additional_charge", func(p broker.Event) error {
		logger.Infof("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	}, broker.Queue("order_capture_additional_charge"))
	if err != nil {
		logger.Error("error", err)
	}

	// Register Handler
	_ = micro_services.RegisterMovieRepositoryHandler(service.Server(), &handler.BuildMovie{
		Broker: defaultBroker,
	})

	//go sub(defaultBroker)

	//go func() {
	//	// Register Struct as Subscriber
	//	e2 := micro.RegisterSubscriber("shippy.cli.consignment", service.Server(), new(subscriber.BuildMovie))
	//
	//	// Register Function as Subscriber
	//	e3 := micro.RegisterSubscriber("shippy.cli.consignment", service.Server(), subscriber.Handler)
	//
	//	println(e2, e3)
	//}()

	//pub1 := micro.NewEvent("shippy.cli.consignment", service.Client())
	//
	//// pub to topic 1
	//go sendEv("shippy.cli.consignment", pub1)

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
