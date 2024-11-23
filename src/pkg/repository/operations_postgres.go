package repository

import (
	"github.com/jmoiron/sqlx"
	models "github.com/kossadda/wallet-service"
)

type OperationPostgres struct {
	db *sqlx.DB
}

func NewOperationPostgres(db *sqlx.DB) *OperationPostgres {
	return &OperationPostgres{
		db: db,
	}
}

func (s *OperationPostgres) Deposit(req models.Request) (int, error) {
	query := "INSERT INTO wallets (wallet_id, operation_type, amount) VALUES ($1, $2, $3) RETURNING wallet_id"
	row := s.db.QueryRow(query, req.WalletID, req.OperationType, req.Amount)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
