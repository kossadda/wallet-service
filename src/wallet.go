package models

import "github.com/google/uuid"

type WalletRequest struct {
	WalletID      uuid.UUID `json:"walletId"`
	OperationType string    `json:"operationType"`
	Amount        int       `json:"amount"`
}
