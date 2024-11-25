package service

import (
	models "github.com/kossadda/wallet-service"
	"github.com/kossadda/wallet-service/pkg/repository"
)

// Operators is an interface that defines the methods for wallet operations.
type Operators interface {
	BalanceChange(req models.Request) (string, error)
	BalanceCheck(id string) (float64, error)
}

// Service is a struct that holds the implementation of the Operators interface.
type Service struct {
	Operators
}

// New creates a new service with the given repository.
// Returns: *Service
func New(repos *repository.Repository) *Service {
	return &Service{
		Operators: NewOperationService(*repos),
	}
}
