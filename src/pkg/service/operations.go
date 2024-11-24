package service

import (
	models "github.com/kossadda/wallet-service"
	"github.com/kossadda/wallet-service/pkg/repository"
)

type OperationService struct {
	repo repository.Repository
}

func NewOperationService(repo repository.Repository) *OperationService {
	return &OperationService{
		repo: repo,
	}
}

func (s *OperationService) BalanceChange(req models.Request) (string, error) {
	return s.repo.Operation.BalanceChange(req)
}

func (s *OperationService) BalanceCheck(id string) (float64, error) {
	return s.repo.Operation.BalanceCheck(id)
}
