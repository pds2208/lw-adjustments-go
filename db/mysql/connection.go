package mysql

import (
	"github.com/rs/zerolog/log"
	"lw-adjustments/config"
	"time"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

type Connection struct {
	DB sqlbuilder.Database
}

func (s *Connection) Connect() error {

	var settings = mysql.ConnectionURL{
		Database: config.Config.Database.Database,
		Host:     config.Config.Database.Server,
		User:     config.Config.Database.User,
		Password: config.Config.Database.Password,
	}

	log.Debug().
		Str("db", config.Config.Database.Database).
		Msg("Connecting to database,")

	sess, err := mysql.Open(settings)

	if err != nil {
		log.Error().
			Err(err).
			Str("db", config.Config.Database.Database).
			Msg("Cannot connect to database")
		return err
	}

	log.Debug().
		Str("db", config.Config.Database.Database).
		Msg("Database Connected")

	if config.Config.Database.Verbose {
		sess.SetLogging(true)
	}

	s.DB = sess

	poolSize := config.Config.ConnectionPool.MaxPoolSize
	maxIdle := config.Config.ConnectionPool.MaxIdleConnections
	maxLifetime := config.Config.ConnectionPool.MaxLifetimeSeconds

	if maxLifetime > 0 {
		maxLifetime = maxLifetime * time.Second
		sess.SetConnMaxLifetime(maxLifetime)
	}

	log.Debug().
		Int("MaxPoolSize", poolSize).
		Int("MaxIdleConnections", maxIdle).
		Dur("MaxLifetime", maxLifetime*time.Second).
		Msg("Connection Attributes:")

	sess.SetMaxOpenConns(poolSize)
	sess.SetMaxIdleConns(maxIdle)

	return nil
}

func (s Connection) Close() {
	if s.DB != nil {
		_ = s.DB.Close()
	}
}
