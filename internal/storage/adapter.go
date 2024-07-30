package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/photo-pixels/platform/log"
)

// Adapter адаптер для работы с базой
type Adapter struct {
	logger log.Logger
	transactor
}

// NewStorageAdapter новый адаптер
func NewStorageAdapter(
	logger log.Logger,
	pool *pgxpool.Pool,
) *Adapter {
	return &Adapter{
		logger:     logger.Named("storage_adapter"),
		transactor: newTransactor(pool),
	}
}
