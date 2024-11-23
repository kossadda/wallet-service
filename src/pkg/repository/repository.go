package repository

import (
	"github.com/jmoiron/sqlx"
	models "github.com/kossadda/wallet-service"
)

type Operation interface {
	Deposit(req models.Request) (int, error)
}

type Repository struct {
	Operation
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Operation: NewOperationPostgres(db),
	}
}
