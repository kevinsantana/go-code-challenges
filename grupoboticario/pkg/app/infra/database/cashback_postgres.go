package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

type PostgresAdapter struct {
	Conn *sql.DB
}

func (P PostgresAdapter) Start() error {
	_, err := P.Conn.Begin()
	if err != nil {
		return fmt.Errorf("error to start postgres transaction: %w", err)
	}

	return nil
}

func (P PostgresAdapter) PurchaseInsert(ctx context.Context, purchase_id string, amount float64, date time.Time) error {
	return nil

}

func (P PostgresAdapter) CreateUser(ctx context.Context, name string, cpf string, email string, password string) (int64, error) {
	var id int64

	err := P.Conn.QueryRowContext(
		ctx,
		`INSERT INTO retailer (NAME_RETAILER, CPF, EMAIL, PSW)
					VALUES ($1, $2, $3, $4) RETURNING id`,
		name,
		cpf,
		email,
		password,
	).Scan(&id)

	if err != nil {
		log.WithFields(map[string]interface{}{
			"NAME_RETAILER": name,
			"CPF":           cpf,
			"EMAIL":         email,
			"PSW":           password,
		}).Error("Error to create new retailer")

		return 0, err
	}

	return id, nil

}
