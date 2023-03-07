package auth

import (
	"context"
	"errors"
	
	log "github.com/sirupsen/logrus"

	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app/models"
)

var (
	ErrForbidden = errors.New("forbidden")
)

type AuthService struct {
	userGateway  UserGateway
}

func NewAuthService(
	userGateway UserGateway,
) *AuthService {
	return &AuthService{
		userGateway:  userGateway,
	}
}

func (service *AuthService) SignerUp(ctx context.Context, requestSignUp models.SignUp) (int64, error) {
	retailerId, err := service.userGateway.SignUp(ctx, requestSignUp)
	
	if err != nil {
		log.WithError(err).Error("Error to sign up user")

		return 0, err
	}

	return retailerId, nil
}