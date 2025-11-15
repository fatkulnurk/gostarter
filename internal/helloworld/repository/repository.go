package repository

import (
	"database/sql"
	"github.com/fatkulnurk/gostarter/internal/helloworld/domain"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.Repository {
	return &Repository{db: db}
}
