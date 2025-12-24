package store

import (
	"context"
	"database/sql"
	"smart-url/internal/traffic"
)

type TrafficLog struct {
	ID          int64
	LinkID      int64
	IPAddress   string // FK to Accounts
	UserAgent   string
	Status      traffic.TrafficStatus // VALID, FRAUD, DUPLICATE
	ProcessedAt string
}

type TrafficLogsStore struct {
	db *sql.DB
}

func (s *TrafficLogsStore) Create(ctx context.Context) error {
	return nil
}
