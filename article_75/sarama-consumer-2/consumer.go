package main

import (
	"github.com/Shopify/sarama"
)

// BrokerConfiguration represents configuration of Kafka brokers and topics
type BrokerConfiguration struct {
	Address string `mapstructure:"address" toml:"address"`
	Topic   string `mapstructure:"topic"   toml:"topic"`
	Group   string `mapstructure:"group"   toml:"group"`
}

// Consumer represents any consumer
type Consumer interface {
	Serve()
	ProcessMessage(msg *sarama.ConsumerMessage) error
	Close() error
}
