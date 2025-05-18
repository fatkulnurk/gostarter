package repository

import (
	"database/sql"
	"github.com/fatkulnurk/gostarter/internal/example/domain"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.IRepository {
	return &Repository{db: db}
}
