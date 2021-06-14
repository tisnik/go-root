package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Message struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Started")

	brokerConfiguration := BrokerConfiguration{
		Address: "localhost:9092",
		Topic:   "test-topic2",
	}

	producer, err := NewKafkaProducer(brokerConfiguration)
	if err != nil {
		log.Error().Err(err).Msg("Can not connect to Kafka")
		return
	}

	log.Info().Str("address", brokerConfiguration.Address).Msg("Connected to Kafka")

	defer producer.Close()

	message := Message{
		ID:      42,
		Name:    "Václav",
		Surname: "Trachta",
	}

	_, _, err = producer.ProduceMessage(message)
	if err != nil {
		log.Error().Err(err).Msg("Unable to produce message")
		return
	}

	log.Info().Msg("Finished")
}
