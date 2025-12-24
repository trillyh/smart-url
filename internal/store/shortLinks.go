package store

import (
	"context"
	"database/sql"
)

type ShortLinks struct {
	ID             int64
	OfferID        int64 // NULL if for personal usage.
	PartnerID      int64 // FK to Accounts
	ShortCode      int64
	DestinationUrl string
}

type ShortLinksStorage struct {
	db *sql.DB
}

func (s *ShortLinksStorage) Create(ctx context.Context) error {
	return nil
}
