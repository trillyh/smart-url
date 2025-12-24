package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Urls interface {
		Create(context.Context) error
	}

	Users interface {
		Create(context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &AccountsStorage{db},
		Urls:  &ShortLinksStorage{db},
	}
}
