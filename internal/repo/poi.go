package repo

import (
	"context"
	"github.com/iter-x/iter-x/internal/repo/ent"
	"github.com/iter-x/iter-x/internal/repo/ent/pointsofinterest"
	"go.uber.org/zap"
)

type PointsOfInterest struct {
	*Tx
	cli    *ent.Client
	logger *zap.SugaredLogger
}

func NewPointsOfInterest(cli *ent.Client, logger *zap.SugaredLogger) *PointsOfInterest {
	return &PointsOfInterest{
		Tx:     &Tx{cli: cli},
		cli:    cli,
		logger: logger.Named("repo.poi"),
	}
}

func (r *PointsOfInterest) SearchPointsOfInterest(ctx context.Context, keyword string, limit int, tx ...*ent.Tx) ([]*ent.PointsOfInterest, error) {
	cli := r.cli.PointsOfInterest
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].PointsOfInterest
	}

	return cli.Query().
		Where(pointsofinterest.NameContains(keyword)).
		WithContinent().
		WithCountry().
		WithState().
		WithCity().
		Limit(limit).
		All(ctx)
}
