package service

import (
	models "github.com/kossadda/wallet-service"
	"github.com/kossadda/wallet-service/pkg/repository"
)

type Operators interface {
	BalanceChange(req models.Request) (string, error)
	BalanceCheck(id string) (float64, error)
}

type Service struct {
	Operators
}

func New(repos *repository.Repository) *Service {
	return &Service{
		Operators: NewOperationService(*repos),
	}
}
