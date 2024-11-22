package repository

type Operation interface{}

type Repository struct {
	Operation
}

func New() *Repository {
	var r Repository

	return &r
}
