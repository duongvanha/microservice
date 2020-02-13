package rabbitmq

import "microservice/pkg/broker"

type RBroker struct {
}

func (r *RBroker) NewBroker() broker.Broker {
	panic("implement me")
}

func (r *RBroker) Init() error {
	panic("implement me")
}

func (r *RBroker) Connect() error {
	panic("implement me")
}

func (r *RBroker) Disconnect() error {
	panic("implement me")
}

func (r *RBroker) Publish(topic string, m *broker.Message) error {
	panic("implement me")
}

func (r *RBroker) Subscribe(topic string, h broker.Handler) error {
	panic("implement me")
}
