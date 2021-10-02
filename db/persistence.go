package db

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"lw-adjustments/config"
	"lw-adjustments/db/mysql"
	"lw-adjustments/types"
	"sync"
)

var cachedConnection Persistence
var connectionMux = &sync.Mutex{}

type Persistence interface {
	Connect() error
	Close()
	GetAdjustments() ([]types.Adjustments, error)
	GetAllAdjustments() ([]types.Adjustments, error)
	DeleteAdjustment(id int) error
}

func GetDefaultPersistenceImpl() (Persistence, error) {
	connectionMux.Lock()
	defer connectionMux.Unlock()

	if cachedConnection != nil {
		log.Debug().
			Str("database", config.Config.Database.Database).
			Msg("Returning cached database connection")
		return cachedConnection, nil
	}

	cachedConnection = &mysql.Connection{}

	if err := cachedConnection.Connect(); err != nil {
		log.Info().
			Err(err).
			Str("database", config.Config.Database.Database).
			Msg("Cannot connect to database")
		cachedConnection = nil
		return nil, fmt.Errorf("cannot connect to database")
	}

	return cachedConnection, nil
}
