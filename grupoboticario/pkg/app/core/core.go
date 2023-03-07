package core

import (
	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app/infra/database"
)

type Service struct {
	Cashback database.CashbackDBGateway
}