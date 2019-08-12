package main

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	server  string = "127.0.0.1:9092"
	groupId string = "haduong"
	topic   string = "two_topic"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          groupId,
	})

	if err != nil {
		panic(err)
	}

	adminClient, err := kafka.NewAdminClientFromConsumer(c)

	if err != nil {
		panic(err)
	}

	_, _ = adminClient.CreateTopics(context.Background(), []kafka.TopicSpecification{{
		Topic:             topic,
		NumPartitions:     3,
		ReplicationFactor: 1,
	}})

	_ = c.SubscribeTopics([]string{topic}, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	_ = c.Close()
}
