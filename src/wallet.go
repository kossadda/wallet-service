package models

import "github.com/google/uuid"

func NewRequest() *request {
	var r request

	return &r
}

type request struct {
	WalletID      uuid.UUID `json:"walletId"`
	OperationType string    `json:"operationType"`
	Amount        int       `json:"amount"`
}
