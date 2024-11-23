package repository

import "github.com/jmoiron/sqlx"

type Operation interface{}

type Repository struct {
	Operation
}

func New(db *sqlx.DB) *Repository {
	return &Repository{}
}
