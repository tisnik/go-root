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
		Group:   "grpX",
	}

	consumer, err := NewKafkaConsumer(brokerConfiguration)
	if err != nil {
		log.Error().Err(err).Msg("Can not connect to Kafka")
		return
	}

	log.Info().Str("address", brokerConfiguration.Address).Msg("Connected to Kafka")

	// zajištění uzavření připojení ke Kafce
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Error().Err(err).Msg("Can not properly close consumer")
		}
	}()

	consumer.Serve()

	log.Info().Msg("Finished")
}
