package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"os"
)

func main() {
	consumer, err := kafka.NewConsumer(
		&kafka.ConfigMap{
			"bootstrap.servers": "127.0.0.1:9092",
			"group.id":          "foo",
			"auto.offset.reset": "smallest",
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	topics := []string{"test"}
	err = consumer.SubscribeTopics(topics, nil)

	if err != nil {
		log.Fatal(err)
	}

	for {
		ev := consumer.Poll(1000)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Println(string(e.Value))
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			break
		default:
			continue
		}
	}

	consumer.Close()
}
