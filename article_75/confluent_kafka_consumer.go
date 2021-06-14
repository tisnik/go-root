// Ukázka použití rozhraní pro systém Apache Kafka představovaného knihovnou
// confluent-kafka-go: implementace konzumenta zpráv.

package main

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const (
	server   = "localhost:9092"
	topic    = "upload"
	group_id = "group1"
)

func main() {
	// konstrukce konzumenta
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          group_id,
		"auto.offset.reset": "earliest",
	})

	// kontrola chyby při připojování ke Kafce
	if err != nil {
		panic(err)
	}

	// i konzumenta je nutné na konci uzavřít
	defer consumer.Close()

	// přihlášení konzumenta do zvoleného tématu (či témat)
	consumer.SubscribeTopics([]string{topic}, nil)

	// postupné čtení zpráv, které byly do zvoleného tématu publikovány
	for {
		message, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s %s\n", message.TopicPartition, string(message.Key), string(message.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, message)
		}
	}
}
