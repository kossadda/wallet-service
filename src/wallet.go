package models

import "github.com/google/uuid"

type WalletRequest struct {
	WalletID      uuid.UUID `json:"walletID"`
	OperationType string    `json:"operationType"`
	Amount        int       `json:"amount"`
}
