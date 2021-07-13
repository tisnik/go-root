package main

import (
	"fmt"
	"io"
	"os"

	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const Address = "localhost"
const Port = 8080
const Enabled = true

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Started")

	log.Info().
		Str("address", Address).
		Int("port", Port).
		Bool("enabled", Enabled).Msg("Server settings")

	if Enabled {
		http.HandleFunc("/", handler)
		where := fmt.Sprintf("%s:%d", Address, Port)
		err := http.ListenAndServe(where, nil)
		if err != nil {
			log.Error().Err(err).Msg("Initialize server")
		}
	}
	log.Info().Msg("Finished")
}

func handler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello world!")
}
