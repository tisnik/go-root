package main

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

// KafkaProducer is an implementation of Producer interface
type KafkaProducer struct {
	Configuration BrokerConfiguration
	Producer      sarama.SyncProducer
}

// NewKafkaProducer constructs new implementation of Producer interface
func NewKafkaProducer(brokerConfiguration BrokerConfiguration) (*KafkaProducer, error) {
	producer, err := sarama.NewSyncProducer([]string{brokerConfiguration.Address}, nil)
	if err != nil {
		log.Error().Err(err).Msg("unable to create a new Kafka producer")
		return nil, err
	}

	return &KafkaProducer{
		Configuration: brokerConfiguration,
		Producer:      producer,
	}, nil
}

// ProduceMessage produces message to selected topic. That function returns
// partition ID and offset of new message or an error value in case of any
// problem on broker side.
func (producer *KafkaProducer) ProduceMessage(message Message) (partitionID int32, offset int64, err error) {
	jsonBytes, err := json.Marshal(message)

	if err != nil {
		log.Error().Err(err).Msg("Couldn't turn notification message into valid JSON")
		return -1, -1, err
	}

	// construct message to be produced using the provided payload (message body)
	producerMsg := &sarama.ProducerMessage{
		Topic: producer.Configuration.Topic,
		Value: sarama.ByteEncoder(jsonBytes),
	}

	// try to produce message
	partitionID, offset, err = producer.Producer.SendMessage(producerMsg)
	if err != nil {
		log.Error().Err(err).Msg("failed to produce message to Kafka")
	} else {
		log.Info().Msgf("message sent to partition %d at offset %d\n", partitionID, offset)
	}
	return
}

// Close allow the Sarama producer to be gracefully closed
func (producer *KafkaProducer) Close() error {
	if err := producer.Producer.Close(); err != nil {
		log.Error().Err(err).Msg("unable to close Kafka producer")
		return err
	}

	return nil
}
