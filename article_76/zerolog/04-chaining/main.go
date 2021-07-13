package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const Address = "localhost"
const Port = 8080
const Enabled = false

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Started")

	log.Info().
		Str("address", Address).
		Int("port", Port).
		Bool("enabled", Enabled).Msg("Server settings")

	log.Info().Msg("Finished")
}
