package data

import (
	"github.com/iter-x/iter-x/internal/data/cache"
	_ "github.com/lib/pq"

	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type (
	Tx struct {
		Cli *ent.Client
	}

	Data struct {
		Tx    *Tx
		Cache cache.Cacher
	}
)

// newTx create a new tx
func newTx(cli *ent.Client) *Tx {
	return &Tx{Cli: cli}
}

func NewConnection(c *conf.Data, logger *zap.SugaredLogger) (*Data, func(), error) {
	logger = logger.Named("repo")

	client, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		logger.Error("failed opening connection to postgres: ", err)
		return nil, nil, err
	}
	d := &Data{
		Tx:    newTx(client),
		Cache: cache.NewCache(c),
	}

	cleanup := func() {
		if err = d.Tx.Cli.Close(); err != nil {
			logger.Error("failed closing connection: ", err)
		} else {
			logger.Info("closing the data resources")
		}
		if err = d.Cache.Close(); err != nil {
			logger.Error("failed closing connection: ", err)
		} else {
			logger.Info("closing the data resources")
		}
	}

	return d, cleanup, nil
}

// GetTx This method checks if there is a transaction in the context, and if so returns the client with the transaction
func (t *Tx) GetTx(ctx context.Context) *ent.Client {
	tx, ok := ctx.Value(contextTxKey{}).(*ent.Tx)
	if ok {
		return tx.Client()
	}
	return t.Cli
}

func (t *Tx) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := t.Cli.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
		}
	}()
	err = fn(tx)
	if err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rErr)
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
