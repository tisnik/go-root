package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/Shopify/sarama"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/rs/zerolog/log"
)

// Messages to be displayed on terminal or written into logs
const (
	versionMessage                 = "Spejbl version 1.0"
	authorsMessage                 = "Pavel Tisnovsky, Red Hat Inc."
	connectionToBrokerMessage      = "Connection to broker"
	operationFailedMessage         = "Operation failed"
	notConnectedToBrokerMessage    = "Not connected to broker"
	brokerConnectionSuccessMessage = "Broker connection OK"
)

// Exit codes
const (
	// ExitStatusOK means that the tool finished with success
	ExitStatusOK = iota
	// ExitStatusConsumerError is returned in case of any consumer-related error
	ExitStatusConsumerError
	// ExitStatusKafkaError is returned in case of any Kafka-related error
	ExitStatusKafkaError
	// ExitStatusHTTPServerError is returned in case the HTTP server can not be started
	ExitStatusHTTPServerError
)

// CliFlags represents structure holding all command line arguments and flags.
type CliFlags struct {
	CheckConnectionToKafka bool
	ShowVersion            bool
	ShowAuthors            bool
}

type ConfigStruct struct {
	// Address represents Kafka address
	Address string `mapstructure:"address" toml:"address"`
	// SecurityProtocol represents the security protocol used by the broker
	SecurityProtocol string `mapstructure:"security_protocol" toml:"security_protocol"`
	// 	CertPath is the path to a file containing the certificate to be used with the broker
	CertPath string `mapstructure:"cert_path" toml:"cert_path"`
	// SaslMechanism is the SASL mechanism used for authentication
	SaslMechanism string `mapstructure:"sasl_mechanism" toml:"sasl_mechanism"`
	// SaslUsername is the username used in case of PLAIN mechanism
	SaslUsername string `mapstructure:"sasl_username" toml:"sasl_username"`
	// SaslPassword is the password used in case of PLAIN mechanism
	SaslPassword string `mapstructure:"sasl_password" toml:"sasl_password"`
	// Topic is name of Kafka topic
	Topic string `mapstructure:"topic" toml:"topic"`
	// Group is name of Kafka group
	Group string `mapstructure:"group" toml:"group"`
	// Enabled is set to true if Kafka consumer is to be enabled
	Enabled bool `mapstructure:"enabled" toml:"enabled"`

	MetricsAddress string
}

// showVersion function displays version information to standard output.
func showVersion() {
	fmt.Println(versionMessage)
}

// showAuthors function displays information about authors to standard output.
func showAuthors() {
	fmt.Println(authorsMessage)
}

// tryToConnectToKafka function just tries to establish connection to Kafka
// broker
func tryToConnectToKafka(configuration *ConfigStruct) (int, error) {
	log.Info().Msg("Checking connection to Kafka")

	// display basic info about broker that will be used
	log.Info().
		Str("broker address", configuration.Address).
		Msg("Broker address")

	// create new broker instance (w/o any checks)
	broker := sarama.NewBroker(configuration.Address)

	// check broker connection
	err := broker.Open(nil)
	if err != nil {
		log.Error().Err(err).Msg(connectionToBrokerMessage)
		return ExitStatusKafkaError, err
	}

	// check if connection remain
	connected, err := broker.Connected()
	if err != nil {
		log.Error().Err(err).Msg(connectionToBrokerMessage)
		return ExitStatusKafkaError, err
	}
	if !connected {
		log.Error().Err(err).Msg(notConnectedToBrokerMessage)
		return ExitStatusConsumerError, err
	}

	// connection was established
	log.Info().Msg(brokerConnectionSuccessMessage)

	// everything seems to be ok
	return ExitStatusOK, nil
}

// startService function tries to start the notification writer service,
// connect to storage and initialize connection to message broker.
func startService(configuration *ConfigStruct) (int, error) {
	// prepare HTTP server with metrics exposed
	err := startHTTPServer(configuration.MetricsAddress)
	if err != nil {
		log.Error().Err(err)
		return ExitStatusHTTPServerError, err
	}

	return ExitStatusOK, nil
}

// startHTTP server starts HTTP or HTTPS server with exposed metrics.
func startHTTPServer(address string) error {
	// setup handlers
	http.Handle("/metrics", promhttp.Handler())

	// start the server
	log.Info().Str("HTTP server address", address).Msg("Starting HTTP server")
	err := http.ListenAndServe(address, nil) // #nosec G114
	if err != nil {
		log.Error().Err(err).Msg("Listen and serve")
		return err
	}
	return nil
}

func doSelectedOperation(configuration *ConfigStruct, cliFlags CliFlags) (int, error) {
	switch {
	case cliFlags.ShowVersion:
		showVersion()
		return ExitStatusOK, nil
	case cliFlags.ShowAuthors:
		showAuthors()
		return ExitStatusOK, nil
	case cliFlags.CheckConnectionToKafka:
		return tryToConnectToKafka(configuration)
	default:
		exitCode, err := startService(configuration)
		return exitCode, err
	}
	// this can not happen: return ExitStatusOK, nil
}

// main function is entry point to the Notification writer service.
func main() {
	var cliFlags CliFlags

	// define and then parse all command line options
	flag.BoolVar(&cliFlags.CheckConnectionToKafka, "check-kafka", false, "check connection to Kafka")
	flag.BoolVar(&cliFlags.ShowVersion, "version", false, "show version")
	flag.BoolVar(&cliFlags.ShowAuthors, "authors", false, "show authors")
	flag.Parse()

	configuration := ConfigStruct{}

	// perform selected operation
	exitStatus, err := doSelectedOperation(&configuration, cliFlags)
	if err != nil {
		log.Err(err).Msg("Do selected operation")
		os.Exit(exitStatus)
		return
	}

	log.Debug().Msg("Finished")
}
