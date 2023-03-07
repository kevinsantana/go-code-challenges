package auth

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app/infra/database"
	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app/models"
)

type AuthorizeAdapter struct {
	cashback database.CashbackDBGateway
}

func NewAuthorizeAdapter(cashback database.CashbackDBGateway) AuthorizeAdapter {
	return AuthorizeAdapter{
		cashback: cashback,
	}
}

func (adapter AuthorizeAdapter) SignUp(ctx context.Context, signUp models.SignUp) (int64, error) {
	parsedPsw := fmt.Sprintf("%s%s%s", "`", signUp.Password, "`")
	h := sha1.New()
	h.Write([]byte(parsedPsw))
	hashedPsw := hex.EncodeToString(h.Sum(nil))

	id, err := adapter.cashback.CreateUser(ctx, signUp.Name, signUp.CPF, signUp.Email, hashedPsw)

	if err != nil {
		log.WithError(err).Error("Error to create user")
	}

	return id, nil

}
