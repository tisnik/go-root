// Ukázka použití rozhraní pro systém Apache Kafka představovaného knihovnou
// Sarama: implementace producenta zpráv.

package main

import (
	"log"

	"github.com/Shopify/sarama"
)

const (
	// KafkaConnectionString obsahuje jméno počítače a port, na kterém běží Kafka broker
	KafkaConnectionString = "localhost:9092"

	// KafkaTopic obsahuje jméno tématu
	KafkaTopic = "test-topic"
)

func main() {
	// konstrukce konzumenta
	producer, err := sarama.NewSyncProducer([]string{KafkaConnectionString}, nil)

	// kontrola chyby při připojování ke Kafce
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to %s", KafkaConnectionString)

	// zajištění uzavření připojení ke Kafce
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// poslání (produkce) zprávy
	msg := &sarama.ProducerMessage{Topic: KafkaTopic, Value: sarama.StringEncoder("testing 123")}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}

	log.Print("Done")
}
