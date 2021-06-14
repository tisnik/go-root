package main

// BrokerConfiguration represents configuration of Kafka brokers and topics
type BrokerConfiguration struct {
	Address string `mapstructure:"address" toml:"address"`
	Topic   string `mapstructure:"topic"   toml:"topic"`
}

// Producer represents any producer
type Producer interface {
	ProduceMessage(message Message) (int32, int64, error)
	Close() error
}
