package repository

import (
	"database/sql"
	"magicauth/internal/magiclink/domain"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.IRepository {
	return &Repository{db: db}
}
