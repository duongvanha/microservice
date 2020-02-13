package broker

type Broker interface {
	NewBroker() Broker
	Init() error
	Connect() error
	Disconnect() error

	Publish(topic string, m *Message) error
	Subscribe(topic string, h Handler) error
}

type Handler func(interface{}) error

type Message struct {
	Header map[string]string
	Body   []byte
}
