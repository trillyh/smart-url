package store

import (
	"context"
	"database/sql"
)

type Accounts struct {
	ID            int64
	Email         string
	password_hash string
	Role          int8
	CreatedAt     string
}

type AccountsStorage struct {
	db *sql.DB
}

func (s *AccountsStorage) Create(ctx context.Context) error {
	return nil
}
