package config

import (
	"context"

	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app/modules/auth"
	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app/infra/database"
	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app/environment"
)


var (
	AuthService *auth.AuthService
)

func Init(ctx context.Context, conf *environment.Config) error {
	database.ConnectPostgres(ctx, conf)

	if conf.RunMigration {
		database.RunMigrations(ctx, conf)
	}

	cashbackDBAdapter := database.PostgresAdapter{
		Conn: database.ConnPostgres,
	}

	AuthService = auth.NewAuthService(
		auth.NewAuthorizeAdapter(cashbackDBAdapter),
	)

	return nil

}
