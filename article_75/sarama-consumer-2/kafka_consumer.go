package main

import (
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

// KafkaConsumer in an implementation of Consumer interface
type KafkaConsumer struct {
	Configuration                        BrokerConfiguration
	ConsumerGroup                        sarama.ConsumerGroup
	numberOfSuccessfullyConsumedMessages uint64
	numberOfErrorsConsumingMessages      uint64
	ready                                chan bool
	cancel                               context.CancelFunc
}

// DefaultSaramaConfig is a config which will be used by default
// here you can use specific version of a protocol for example
// useful for testing
var DefaultSaramaConfig *sarama.Config

// NewKafkaConsumer constructs new implementation of Consumer interface
func NewKafkaConsumer(brokerCfg BrokerConfiguration) (*KafkaConsumer, error) {
	return NewWithSaramaConfig(brokerCfg, DefaultSaramaConfig)
}

// NewWithSaramaConfig constructs new implementation of Consumer interface with custom sarama config
func NewWithSaramaConfig(
	brokerConfiguration BrokerConfiguration,
	saramaConfig *sarama.Config,
) (*KafkaConsumer, error) {
	if saramaConfig == nil {
		saramaConfig = sarama.NewConfig()
		saramaConfig.Version = sarama.V0_10_2_0

		/* TODO: we need to do it in production code
		if brokerCfg.Timeout > 0 {
			saramaConfig.Net.DialTimeout = brokerCfg.Timeout
			saramaConfig.Net.ReadTimeout = brokerCfg.Timeout
			saramaConfig.Net.WriteTimeout = brokerCfg.Timeout
		}
		*/
	}

	consumerGroup, err := sarama.NewConsumerGroup([]string{brokerConfiguration.Address}, brokerConfiguration.Group, saramaConfig)
	if err != nil {
		return nil, err
	}

	consumer := &KafkaConsumer{
		Configuration:                        brokerConfiguration,
		ConsumerGroup:                        consumerGroup,
		numberOfSuccessfullyConsumedMessages: 0,
		numberOfErrorsConsumingMessages:      0,
		ready:                                make(chan bool),
	}

	return consumer, nil
}

// Serve starts listening for messages and processing them. It blocks current thread.
func (consumer *KafkaConsumer) Serve() {
	ctx, cancel := context.WithCancel(context.Background())
	consumer.cancel = cancel

	go func() {
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := consumer.ConsumerGroup.Consume(ctx, []string{consumer.Configuration.Topic}, consumer); err != nil {
				log.Fatal().Err(err).Msg("Unable to recreate Kafka session")
			}

			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				log.Info().Err(ctx.Err()).Msg("Stopping consumer")
				return
			}

			log.Info().Msg("Created new kafka session")

			consumer.ready = make(chan bool)
		}
	}()

	// Await till the consumer has been set up
	log.Info().Msg("Waiting for consumer to become ready")
	<-consumer.ready
	log.Info().Msg("Finished waiting for consumer to become ready")

	// Actual processing is done in goroutine created by sarama (see ConsumeClaim below)
	log.Info().Msg("Started serving consumer")
	<-ctx.Done()
	log.Info().Msg("Context cancelled, exiting")

	cancel()
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *KafkaConsumer) Setup(sarama.ConsumerGroupSession) error {
	log.Info().Msg("New session has been setup")
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *KafkaConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	log.Info().Msg("New session has been finished")
	return nil
}

// ConsumeClaim starts a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *KafkaConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Info().
		Int64("offset", claim.InitialOffset()).
		Msg("Starting messages loop")

	for message := range claim.Messages() {
		consumer.HandleMessage(message)

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
			log.Error().
				Err(err).
				Msg("Unable to close consumer group")
		}
	}

	return nil
}

// GetNumberOfSuccessfullyConsumedMessages returns number of consumed messages
// since creating KafkaConsumer obj
func (consumer *KafkaConsumer) GetNumberOfSuccessfullyConsumedMessages() uint64 {
	return consumer.numberOfSuccessfullyConsumedMessages
}

// GetNumberOfErrorsConsumingMessages returns number of errors during consuming messages
// since creating KafkaConsumer obj
func (consumer *KafkaConsumer) GetNumberOfErrorsConsumingMessages() uint64 {
	return consumer.numberOfErrorsConsumingMessages
}

// HandleMessage handles the message and does all logging, metrics, etc
func (consumer *KafkaConsumer) HandleMessage(msg *sarama.ConsumerMessage) {
	log.Info().
		Int64("offset", msg.Offset).
		Int32("partition", msg.Partition).
		Str("topic", msg.Topic).
		Msg("Started processing message")

	err := consumer.ProcessMessage(msg)

	log.Info().
		Int64("offset", msg.Offset).
		Int32("partition", msg.Partition).
		Str("topic", msg.Topic).
		Msg("Finished processing message")

	// Something went wrong while processing the message.
	if err != nil {
		log.Error().
			Err(err).
			Msg("Error processing message consumed from Kafka")
		consumer.numberOfErrorsConsumingMessages++
	} else {
		// The message was processed successfully.
		consumer.numberOfSuccessfullyConsumedMessages++
	}
}

// ProcessMessage processes an incoming message
func (consumer *KafkaConsumer) ProcessMessage(msg *sarama.ConsumerMessage) error {
	log.Info().
		Str("key", string(msg.Key)).
		Str("value", string(msg.Value)).
		Msg("Consumed")

	var deserialized Message

	err := json.Unmarshal(msg.Value, &deserialized)
	if err != nil {
		return err
	}

	log.Info().Int("ID", deserialized.ID).
		Str("Name", deserialized.Name).
		Str("Surname", deserialized.Surname).
		Msg("Deserialized message")

	return nil
}
