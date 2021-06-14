// Ukázka použití rozhraní pro systém Apache Kafka představovaného knihovnou
// Sarama: implementace konzumenta zpráv.

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
	consumer, err := sarama.NewConsumer([]string{KafkaConnectionString}, nil)

	// kontrola chyby při připojování ke Kafce
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to %s", KafkaConnectionString)

	// zajištění uzavření připojení ke Kafce
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// přihlášení ke zvolenému tématu
	partitionConsumer, err := consumer.ConsumePartition(KafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal(err)
	}

	// zajištění ukončení přihlášení ke zvolenému tématu
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// postupné čtení zpráv, které byly do zvoleného tématu publikovány
	consumed := 0
	for {
		msg := <-partitionConsumer.Messages()
		// vypíšeme pouze offset zprávy, její klíč a tělo (value, payload)
		log.Printf("Consumed message offset %d: %s:%s", msg.Offset, msg.Key, msg.Value)
		consumed++
	}

	// výpis počtu zpracovaných zpráv (ovšem sem se stejně nedostaneme :-)
	log.Printf("Consumed: %d", consumed)
	log.Print("Done")
}
