package repository

import (
	"database/sql"
	"fmt"

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

func (s *OperationPostgres) BalanceChange(req models.Request) (string, error) {
	var id string
	var balance float64
	var err error

	tx, err := s.db.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	query := "SELECT wallet_id, balance FROM wallets WHERE wallet_id = $1 FOR UPDATE"
	row := tx.QueryRow(query, req.WalletID)
	if err := row.Scan(&id, &balance); err != nil {
		if err == sql.ErrNoRows {
			if req.OperationType == "DEPOSIT" {
				query = "INSERT INTO wallets (wallet_id, balance) VALUES ($1, $2)"
				_, err = tx.Exec(query, req.WalletID, req.Amount)
				if err != nil {
					return "", err
				}
			} else {
				return "", fmt.Errorf("you cannot withdraw from a non-existent account")
			}
		} else {
			return "", err
		}
	} else {
		if req.OperationType == "DEPOSIT" {
			balance += req.Amount
		} else if req.OperationType == "WITHDRAW" {
			if balance < req.Amount {
				return "", fmt.Errorf("insufficient funds")
			}
			balance -= req.Amount
		} else {
			return "", fmt.Errorf("invalid operation type")
		}

		query = "UPDATE wallets SET balance = $1 WHERE wallet_id = $2"
		_, err = tx.Exec(query, balance, req.WalletID)
		if err != nil {
			return "", err
		}
	}

	if err = tx.Commit(); err != nil {
		return "", err
	}

	return id, nil
}

func (s *OperationPostgres) BalanceCheck(id string) (float64, error) {
	var balance float64

	tx, err := s.db.Begin()
	if err != nil {
		return 0.0, err
	}

	query := "SELECT balance FROM wallets WHERE wallet_id = $1"
	row := tx.QueryRow(query, id)
	if err := row.Scan(&balance); err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			return 0.0, fmt.Errorf("no wallet with ID %s", id)
		}
		return 0.0, err
	}

	if err = tx.Commit(); err != nil {
		return 0.0, err
	}

	return balance, nil
}
