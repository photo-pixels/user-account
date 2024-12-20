package pgrepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// PgConfig конфиг базы
type PgConfig struct {
	ConnString string `yaml:"conn_string"`
}

// NewPgConn новый пул коннектов зп
func NewPgConn(ctx context.Context, cfg PgConfig) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(ctx, cfg.ConnString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	return conn, nil
}
