package main

import (
	"github.com/rs/zerolog/log"
)

const Address = "localhost"
const Port = 8080
const Enabled = false

func main() {
	log.Info().Msg("Started")

	log.Info().Str("address", Address).Msg("Server address")
	log.Info().Int("port", Port).Msg("Server port")
	log.Info().Bool("enabled", Enabled).Msg("Server enabled")

	log.Info().Msg("Finished")
}
