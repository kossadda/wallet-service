package repository

import (
	"github.com/jmoiron/sqlx"

	models "github.com/kossadda/wallet-service"
)

// Operation is an interface that defines the methods for wallet operations.
type Operation interface {
	BalanceChange(req models.Request) (string, error)
	BalanceCheck(id string) (float64, error)
}

// Repository is a struct that holds the implementation of the Operation interface.
type Repository struct {
	Operation
}

// New creates a new repository with the given database connection.
// Returns: *Repository
func New(db *sqlx.DB) *Repository {
	return &Repository{
		Operation: NewOperationPostgres(db),
	}
}
