package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"lw-adjustments/config"
	"lw-adjustments/db"
	"lw-adjustments/services"
	"net/http"
	"os"
	"time"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if config.Config.LogFormat == "Terminal" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true, TimeFormat: time.RFC3339})
	}

	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.Use(logger.SetLogger())

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

	listenAddress := config.Config.Service.ListenAddress

	writeTimeout, err := time.ParseDuration(config.Config.Service.WriteTimeout)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "LW").
			Msgf("writeTimeout configuration error")
	}

	readTimeout, err := time.ParseDuration(config.Config.Service.ReadTimeout)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "LW").
			Msgf("readTimeout configuration error")
	}

	router.Use(cors.Default())

	srv := &http.Server{
		Handler:      router,
		Addr:         listenAddress,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}

	go func() {
		// loop forever
		services.NewAdjustmentService(dbase).SyncAdjustments()
	}()

	log.Info().
		Str("listenAddress", listenAddress).
		Str("writeTimeout", writeTimeout.String()).
		Str("readTimeout", readTimeout.String()).
		Msg("Lewis & Wood Adjustments: Waiting for requests")

	err = srv.ListenAndServe()
	log.Fatal().
		Err(err).
		Str("service", "LW").
		Msgf("ListenAndServe failed")
}
