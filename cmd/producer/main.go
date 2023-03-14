package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"os"
	"time"
)

func main() {
	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": "127.0.0.1:9092",
			"acks":              "all",
		},
	)

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	topicName := "test"

	deliveryChan := make(chan kafka.Event, 10000)

	for {
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topicName, Partition: kafka.PartitionAny},
			Value:          []byte(fmt.Sprintf("hello, now %s", time.Now().String()))},
			deliveryChan,
		)

		e := <-deliveryChan
		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
		time.Sleep(500 * time.Millisecond)
	}
	close(deliveryChan)
}
