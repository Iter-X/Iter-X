package repo

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/repo/ent"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Tx struct {
	cli *ent.Client
}

var ProviderSet = wire.NewSet(NewConnection, NewAuth, NewTrip, NewPointsOfInterest, NewTransactionRepository)

func NewConnection(c *conf.Data, logger *zap.SugaredLogger) (*ent.Client, func(), error) {
	logger = logger.Named("repo")

	client, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		logger.Error("failed opening connection to postgres: ", err)
		return nil, nil, err
	}

	cleanup := func() {
		err = client.Close()
		if err != nil {
			logger.Error("failed closing connection: ", err)
		} else {
			logger.Info("closing the data resources")
		}
	}

	return client, cleanup, nil
}

func (t *Tx) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := t.cli.Tx(ctx)
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
