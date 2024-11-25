package service

import (
	models "github.com/kossadda/wallet-service"
	"github.com/kossadda/wallet-service/pkg/repository"
)

// OperationService is a struct that handles wallet operations.
type OperationService struct {
	repo repository.Repository
}

// NewOperationService creates a new operation service with the given repository.
// Returns: *OperationService
func NewOperationService(repo repository.Repository) *OperationService {
	return &OperationService{
		repo: repo,
	}
}

// BalanceChange changes the balance of the wallet.
// Returns: string, error
func (s *OperationService) BalanceChange(req models.Request) (string, error) {
	return s.repo.Operation.BalanceChange(req)
}

// BalanceCheck checks the balance of the wallet.
// Returns: float64, error
func (s *OperationService) BalanceCheck(id string) (float64, error) {
	return s.repo.Operation.BalanceCheck(id)
}
