// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam příkladů ze sedmdesáté páté části:
//    https://github.com/tisnik/go-root/blob/master/article_75/README.md

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
