package main

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

// Consumer represents any consumer of insights-rules messages
type Consumer interface {
	Serve()
	Close() error
	HandleMessage(msg *sarama.ConsumerMessage) error
}

type KafkaConsumer struct {
	ConsumerGroup                        sarama.ConsumerGroup
	topic                                string
	numberOfSuccessfullyConsumedMessages uint64
	numberOfErrorsConsumingMessages      uint64
	ready                                chan bool
	cancel                               context.CancelFunc
}

type OCPRulesConsumer struct {
	KafkaConsumer
}

type DVOConsumer struct {
	KafkaConsumer
}

func (consumer *KafkaConsumer) Serve() {
	ctx, cancel := context.WithCancel(context.Background())
	consumer.cancel = cancel

	go func() {
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := consumer.ConsumerGroup.Consume(ctx, []string{consumer.topic}, consumer); err != nil {
				log.Fatal().Err(err).Msg("unable to recreate kafka session")
			}

			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}

			log.Info().Msg("created new kafka session")

			consumer.ready = make(chan bool)
		}
	}()

	// Await till the consumer has been set up
	log.Info().Msg("waiting for consumer to become ready")
	<-consumer.ready
	log.Info().Msg("finished waiting for consumer to become ready")

	// Actual processing is done in goroutine created by sarama (see ConsumeClaim below)
	log.Info().Msg("started serving consumer")
	<-ctx.Done()
	log.Info().Msg("context cancelled, exiting")

	cancel()
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *KafkaConsumer) Setup(sarama.ConsumerGroupSession) error {
	log.Info().Msg("new session has been setup")
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *KafkaConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	log.Info().Msg("new session has been finished")
	return nil
}

// ConsumeClaim starts a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *KafkaConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Info().
		Int64("offset", claim.InitialOffset()).
		Msg("starting messages loop")

	for message := range claim.Messages() {
		err := consumer.HandleMessage(message)
		if err != nil {
			// already handled in HandleMessage, just log
			log.Error().Err(err).Msg("Problem while handling the message")
		}
		session.MarkMessage(message, "")
	}

	return nil
}

// Close method closes all resources used by consumer
func (consumer *KafkaConsumer) Close() error {
	if consumer.cancel != nil {
		consumer.cancel()
	}

	if consumer.ConsumerGroup != nil {
		if err := consumer.ConsumerGroup.Close(); err != nil {
			log.Error().Err(err).Msg("unable to close consumer group")
		}
	}
	return nil
}

// GetNumberOfSuccessfullyConsumedMessages returns number of consumed messages
// since creating OCPRulesConsumer obj
func (consumer *KafkaConsumer) GetNumberOfSuccessfullyConsumedMessages() uint64 {
	return consumer.numberOfSuccessfullyConsumedMessages
}

// GetNumberOfErrorsConsumingMessages returns number of errors during consuming messages
// since creating OCPRulesConsumer obj
func (consumer *KafkaConsumer) GetNumberOfErrorsConsumingMessages() uint64 {
	return consumer.numberOfErrorsConsumingMessages
}

func (consumer *KafkaConsumer) HandleMessage(msg *sarama.ConsumerMessage) error {
	log.Info().Msg("KafkaConsumer: handle message")
	return nil
}

func NewKafkaConsumer(address string, topic string, group string) (*KafkaConsumer, error) {
	log.Print("New Kafka consumer")

	consumerGroup, err := sarama.NewConsumerGroup([]string{address}, group, sarama.NewConfig())
	if err != nil {
		log.Error().Err(err).Msg("Unable to create consumer group")
		return nil, err
	}

	consumer := &KafkaConsumer{
		ConsumerGroup:                        consumerGroup,
		topic:                                topic,
		numberOfSuccessfullyConsumedMessages: 0,
		numberOfErrorsConsumingMessages:      0,
		ready:                                make(chan bool),
	}
	return consumer, nil
}

func NewOCPRulesConsumer(address string, topic string, group string) (*OCPRulesConsumer, error) {
	log.Info().Msg("New OCP rules consumer")

	kafkaConsumer, err := NewKafkaConsumer(address, topic, group)
	if err != nil {
		log.Error().Err(err).Msg("Unable to create Kafka consumer")
		return nil, err
	}

	consumer := &OCPRulesConsumer{
		KafkaConsumer: *kafkaConsumer,
	}
	return consumer, nil
}

func NewDVOConsumer(address string, topic string, group string) (*DVOConsumer, error) {
	log.Info().Msg("New DVO consumer")

	kafkaConsumer, err := NewKafkaConsumer(address, topic, group)
	if err != nil {
		log.Error().Err(err).Msg("Unable to create Kafka consumer")
		return nil, err
	}

	consumer := &DVOConsumer{
		KafkaConsumer: *kafkaConsumer,
	}
	return consumer, nil
}

func (consumer *OCPRulesConsumer) HandleMessage(msg *sarama.ConsumerMessage) error {
	log.Info().Msg("OCPRulesConsumer: handle message")
	return nil
}

func main() {
	log.Info().Msg("Starting")

	consumer1, _ := NewOCPRulesConsumer("localhost:9092", "topic1", "group1")
	println(consumer1)
	consumer1.Serve()

	consumer2, _ := NewDVOConsumer("localhost:9092", "topic2", "group2")
	println(consumer2)

	log.Info().Msg("Finishing")
}
