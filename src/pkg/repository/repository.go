package repository

import (
	"github.com/jmoiron/sqlx"
	models "github.com/kossadda/wallet-service"
)

type Operation interface {
	BalanceChange(req models.Request) (string, error)
	BalanceCheck(id string) (float64, error)
}

type Repository struct {
	Operation
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Operation: NewOperationPostgres(db),
	}
}
