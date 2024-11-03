package main

import (
	"github.com/RIDOS/echo-hello/internal/pkg/app"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	a, err := app.New()
	if err != nil {
		log.Fatal().Err(err).Msg("Application initialization failed")
	}

	err = a.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("Application run failed")
	}
}
