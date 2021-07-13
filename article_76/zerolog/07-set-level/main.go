package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const Address = "localhost"
const Port = 8080
const Enabled = true

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Started")

	log.Trace().
		Str("level", "trace").
		Str("address", Address).
		Int("port", Port).
		Bool("enabled", Enabled).Msg("Server settings")

	log.Debug().
		Str("level", "debug").
		Str("address", Address).
		Int("port", Port).
		Bool("enabled", Enabled).Msg("Server settings")

	log.Info().
		Str("level", "info").
		Str("address", Address).
		Int("port", Port).
		Bool("enabled", Enabled).Msg("Server settings")

	log.Warn().
		Str("level", "warn").
		Str("address", Address).
		Int("port", Port).
		Bool("enabled", Enabled).Msg("Server settings")

	log.Error().
		Str("level", "error").
		Str("address", Address).
		Int("port", Port).
		Bool("enabled", Enabled).Msg("Server settings")

	log.Fatal().
		Str("level", "fatal").
		Str("address", Address).
		Int("port", Port).
		Bool("enabled", Enabled).Msg("Server settings")

	log.Panic().
		Str("level", "panic").
		Str("address", Address).
		Int("port", Port).
		Bool("enabled", Enabled).Msg("Server settings")

	log.Info().Msg("Finished")
}
