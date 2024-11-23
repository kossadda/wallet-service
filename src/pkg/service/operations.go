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

func (s *OperationService) Deposit(req models.Request) (int, error) {
	return s.repo.Operation.Deposit(req)
}
