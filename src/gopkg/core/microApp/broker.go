package micro_app

import (
	microBroker "github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
)

func GetBroker() (broker microBroker.Broker, err error) {
	broker = rabbitmq.NewBroker()

	if err := cmd.Init(); err != nil {
		logger.Fatalf("Cmd Init error: %v", err)
	}

	if err := broker.Init(
		microBroker.Addrs("amqp://bkgo:bkgopass@rabbitmq:5672/vhost"),
		rabbitmq.ExchangeName("post_order_processor"),
	); err != nil {
		logger.Errorf("Broker Init error: %v", err)
		return
	}

	if err := broker.Connect(); err != nil {
		logger.Errorf("Broker Connect error: %v", err)
		return
	}

	return
}
