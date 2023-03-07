package database

import (
	"context"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	log "github.com/sirupsen/logrus"

	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app/environment"
)

func RunMigrations(ctx context.Context, conf *environment.Config) {
	PostgresMigrate(ctx, conf)
}

func PostgresMigrate(ctx context.Context, conf *environment.Config) {
	folder := "."

	if f, ok := os.LookupEnv("DATABASE_MIGRATION_ROOT"); ok {
		folder = f
	}

	if conf.RunMigration == true {
		driver, err := postgres.WithInstance(ConnPostgres, &postgres.Config{})
		if err != nil {
			log.WithError(err).Panic("Error on creating migration driver instance")
		}
		m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s/database/migrations", folder), "db_cashback", driver)
		if err != nil {
			log.WithError(err).Panic("Error creating migrator instance")
		}
		
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.WithError(err).Panic("Migration could not be completed")
		}
	}

	

}