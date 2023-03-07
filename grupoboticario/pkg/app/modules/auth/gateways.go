package auth

import (
	"context"

	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app/models"
)

type UserGateway interface {
	SignUp(ctx context.Context, request models.SignUp) (int64, error)
}
