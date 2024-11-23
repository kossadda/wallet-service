package service

import (
	models "github.com/kossadda/wallet-service"
	"github.com/kossadda/wallet-service/pkg/repository"
)

type Operators interface {
	Deposit(req models.Request) (int, error)
}

type Service struct {
	Operators
}

func New(repos *repository.Repository) *Service {
	return &Service{
		Operators: NewOperationService(*repos),
	}
}
