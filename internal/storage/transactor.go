package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/photo-pixels/user-account/internal/storage/db"
)

// Transactor интерфейс транзакций
type Transactor interface {
	RunTransaction(ctx context.Context, txFunc func(ctxTx context.Context) error) error
}

// RunTransaction старт транзакции
func (a *Adapter) RunTransaction(ctx context.Context, txFunc func(ctxTx context.Context) error) error {
	return a.transactor.runTransaction(ctx, txFunc)
}

type contextKey string

const txKey contextKey = "pgx_tx"

type transactor struct {
	pool *pgxpool.Pool
}

func newTransactor(pool *pgxpool.Pool) transactor {
	return transactor{pool: pool}
}

func (t *transactor) getTX(ctx context.Context) db.DBTX {
	if tx, ok := ctx.Value(txKey).(pgx.Tx); ok {
		return tx
	}
	return t.pool
}

func (t *transactor) getQueries(ctx context.Context) *db.Queries {
	return db.New(t.getTX(ctx))
}

func (t *transactor) runTransaction(ctx context.Context, txFunc func(ctxTx context.Context) error) error {
	if _, ok := ctx.Value(txKey).(pgx.Tx); ok {
		return txFunc(ctx)
	}

	tx, err := t.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("con.BeginTx: %w", err)
	}

	txCtx := context.WithValue(ctx, txKey, tx)
	err = txFunc(txCtx)

	if err != nil {
		rollBackErr := tx.Rollback(ctx)
		if rollBackErr != nil {
			return fmt.Errorf("tx.Rollback: %w", rollBackErr)
		}
		return printError(err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("tx.Commit: %w", printError(err))
	}

	return nil
}
