// Ukázka použití rozhraní pro systém Apache Kafka představovaného knihovnou
// confluent-kafka-go: implementace producenta zpráv.

package main

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const (
	server = "localhost"
)

func main() {
	topic := "upload"

	// konstrukce producenta
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": server,
	})

	// kontrola chyby při připojování ke Kafce
	if err != nil {
		panic(err)
	}

	// producenta zpráv je nutné na konci odpojit
	defer producer.Close()

	// funkce volaná pro každou událost, která při práci s Kafkou může nastat
	go func() {
		for event := range producer.Events() {
			switch ev := event.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// vytváření a produkce zpráv posílaných do zvoleného tématu
	for i := 0; i < 100; i++ {
		text := fmt.Sprintf("Message #%d", i)
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(text),
		}, nil)
	}
	producer.Flush(15 * 1000)
}
