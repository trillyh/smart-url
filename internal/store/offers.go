package store

import (
	"context"
	"database/sql"
)

type Offer struct {
	ID              int64
	BrandID         int64 // FK to accounts
	Name            string
	PayoutPerClick  float64
	TotalBudget     float64
	RemainingBudget float64
}

type OffersStore struct {
	db *sql.DB
}

func (s *OffersStore) Create(ctx context.Context) error {
	return nil
}
