package database

import (
	"context"
	"time"
)

type CashbackDBGateway interface {
	PurchaseInsert(ctx context.Context, purchase_id string, amount float64, date time.Time) error
	CreateUser(ctx context.Context, name string, cpf string, email string, password string) (int64, error)
}