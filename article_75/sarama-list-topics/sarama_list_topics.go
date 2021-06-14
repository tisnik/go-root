// Ukázka použití rozhraní pro systém Apache Kafka představovaného knihovnou
// Sarama: výpis informací o tématech.

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
	// konstrukce rozhraní k brokerovi
	broker := sarama.NewBroker(KafkaConnectionString)

	// kontrola chyby při připojování k brokerovi
	err := broker.Open(nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to %s", KafkaConnectionString)

	request := sarama.MetadataRequest{Topics: []string{KafkaTopic}}
	response, err := broker.GetMetadata(&request)
	if err != nil {
		_ = broker.Close()
		log.Fatal(err)
	}

	if len(response.Topics) == 1 {
		log.Print("There is one topic active in the cluster.")
	} else {
		log.Print("There are", len(response.Topics), "topics active in the cluster.")
	}

	if err = broker.Close(); err != nil {
		log.Fatal(err)
	}

	log.Print("Done")
}
