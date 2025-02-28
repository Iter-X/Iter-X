package repo

import (
	"context"

	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/repo/ent"
)

type TransactionRepoImpl struct {
	cli    *ent.Client
	logger *zap.SugaredLogger
}

// contextTxKey The context used to host the transaction
type contextTxKey struct{}

// NewTransactionRepository .
func NewTransactionRepository(cli *ent.Client, logger *zap.SugaredLogger) repository.Transaction {
	return &TransactionRepoImpl{
		cli:    cli,
		logger: logger,
	}
}

func (t *TransactionRepoImpl) Exec(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, ok := ctx.Value(contextTxKey{}).(*ent.Tx)
	if ok {
		return fn(ctx)
	}
	// start transaction
	tx, err := t.cli.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			if err = tx.Rollback(); err != nil {
				t.logger.Errorf("rollback failure from panic recover，%+v", err)
			}
			panic(v)
		}
	}()

	txCtx := context.WithValue(ctx, contextTxKey{}, tx)
	if err = fn(txCtx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			t.logger.Errorf("rollback failure from panic recover, %+v", rollbackErr)
			return rollbackErr
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		t.logger.Error("commit failure", err)
		return err
	}
	return nil
}
