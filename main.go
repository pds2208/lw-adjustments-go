package main

import (
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"lw-adjustments/config"
	"lw-adjustments/db"
	"lw-adjustments/services"
	"os"
	"time"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if config.Config.LogFormat == "Terminal" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true, TimeFormat: time.RFC3339})
	}

	// Command line flag overrides the configuration file
	debug := flag.Bool("debug", false, "sets log level to debug")

	log.Info().Msg("Lewis & Wood Adjustments")

	flag.Parse()

	if *debug || config.Config.LogLevel == "Debug" {
		log.Info().Msg("Debug mode set")
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	dbase, err := db.GetDefaultPersistenceImpl()

	if err != nil {
		log.Error().Err(err)
		os.Exit(1)
	}

	// loop forever
	services.NewAdjustmentService(dbase).SyncAdjustments()
}
