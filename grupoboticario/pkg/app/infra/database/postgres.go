package database

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app/environment"
)

var ConnPostgres *sql.DB

func ConnectPostgres(ctx context.Context, conf *environment.Config) {
	var err error

	ConnPostgres, err = sql.Open("nrpostgres", conf.HostDB)
	if err != nil {
		log.WithError(err).Panic("Error to open database connection")
	}

	ConnPostgres.SetConnMaxLifetime(conf.ConnMaxLifeTime)
	ConnPostgres.SetMaxOpenConns(conf.MaxOpenConns)
	ConnPostgres.SetMaxIdleConns(conf.MaxIdleConns)

	if err := ConnPostgres.Ping(); err != nil {
		log.WithError(err).Panic("Error to connect with database")
	}

	log.Info("Database connected")
}

func ClosePostgres(ctx context.Context) {
	if ConnPostgres == nil {
		return
	}

	if err := ConnPostgres.Close(); err != nil {
		log.WithError(err).Error("Error to close database connection")
	}

	log.Info("Database connection closed")
}