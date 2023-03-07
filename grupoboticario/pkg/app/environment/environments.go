package environment

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Env             string        `envconfig:"ENV" json:"ENV" default:"DEV"`
	Host            string        `envconfig:"HOST" json:"HOST" default:"127.0.0.1"`
	Port            string        `envconfig:"PORT" json:"PORT" default:"5000"`
	ConnMaxLifeTime time.Duration `envconfig:"POSTGRES_CONNMAXLIFETIME" json:"POSTGRES_CONNMAXLIFETIME" default:"5m"`
	MaxOpenConns    int           `envconfig:"POSTGRES_MAXOPENCONNS" json:"POSTGRES_MAXOPENCONNS" default:"25"`
	MaxIdleConns    int           `envconfig:"POSTGRES_MAXIDLECONNS" json:"POSTGRES_MAXIDLECONNS" default:"25"`
	HostDB          string        `envconfig:"POSTGRES_HOST" json:"POSTGRES_HOST" default:"postgres://postgres:secret@localhost:5432/db_cashback?sslmode=disable"`
	RunMigration    bool          `envconfig:"DATABASE_MIGRATE" json:"DATABASE_MIGRATE" default:"false"`
}

var config Config

func LoadConfig(ctx context.Context) *Config {
	conf := &Config{}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&config); err != nil {
		log.WithError(err).Error("Error to unmarshal envs")
	}

	log.WithField("config", conf).Info("Success to load environments")

	return conf
}
