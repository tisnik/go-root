package main

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
)

// Consumer represents any consumer of insights-rules messages
type Consumer interface {
	Serve()
	Close() error
	HandleMessage(msg *sarama.ConsumerMessage) error
}

type KafkaConsumer struct {
	ConsumerGroup                        sarama.ConsumerGroup
	numberOfSuccessfullyConsumedMessages uint64
	numberOfErrorsConsumingMessages      uint64
	ready                                chan bool
	cancel                               context.CancelFunc
}

func (kafkaConsumer KafkaConsumer) start() error {
	log.Print("Kafka Consumer start")
	return nil
}

func (kafkaConsumer KafkaConsumer) stop() error {
	log.Print("Kafka Consumer stop")
	return nil
}

type OCPRulesConsumer struct {
	KafkaConsumer
}

type DVOConsumer struct {
	KafkaConsumer
}

func NewOCPRulesConsumer() OCPRulesConsumer {
	log.Print("New OCP rules consumer")
	consumer := OCPRulesConsumer{}
	return consumer
}

func NewDVOConsumer() DVOConsumer {
	log.Print("New DVO consumer")
	consumer := DVOConsumer{}
	return consumer
}

func main() {
	log.Println("Starting")

	consumer1 := NewOCPRulesConsumer()
	consumer1.start()

	consumer2 := NewDVOConsumer()
	consumer2.start()

	consumer1.stop()
	consumer2.stop()
	log.Println("Finishing")
}
