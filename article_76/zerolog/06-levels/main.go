// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá šestá část:
//    Jazyk Go a textový terminál ve funkci základního prvku uživatelského rozhraní
//    https://www.root.cz/clanky/jazyk-go-a-textovy-terminal-ve-funkci-zakladniho-prvku-uzivatelskeho-rozhrani/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam příkladů ze sedmdesáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_76/README.md

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
