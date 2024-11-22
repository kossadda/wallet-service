package models

import "github.com/google/uuid"

func NewRequest() *request {
	return &request{}
}

type request struct {
	WalletID      uuid.UUID `json:"walletId"`
	OperationType string    `json:"operationType"`
	Amount        int       `json:"amount"`
}
