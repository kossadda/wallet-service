package models

type Request struct {
	// id            int    `json:"-"`
	WalletID      string  `json:"walletId" binding:"required"`
	OperationType string  `json:"operationType" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
}

func NewRequest() *Request {
	return &Request{}
}
