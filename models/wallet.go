package models

// Request represents a wallet operation request.
type Request struct {
	WalletID      string  `json:"walletId" binding:"required"`
	OperationType string  `json:"operationType" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
}

// NewRequest creates a new wallet operation request.
// Returns: *Request
func NewRequest() *Request {
	return &Request{}
}
