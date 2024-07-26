package database

import (
	"context"

	"github.com/andredubov/chat-server/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

// NewPostgresConnection returns an instance for connection to the postgres database
func NewPostgresConnection(cfg config.PostgresConfig) (*pgxpool.Pool, error) {
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, cfg.DSN())

	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
