package service

import "github.com/kossadda/wallet-service/pkg/repository"

type Operation interface{}

type Service struct {
	Operation
}

func New(repos *repository.Repository) *Service {
	return &Service{}
}
