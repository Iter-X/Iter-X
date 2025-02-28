package impl

import (
	"context"

	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/pointsofinterest"
)

type PointsOfInterest struct {
	*data.Tx
	logger *zap.SugaredLogger
}

func NewPointsOfInterest(tx *data.Tx, logger *zap.SugaredLogger) *PointsOfInterest {
	return &PointsOfInterest{
		Tx:     tx,
		logger: logger.Named("repo.poi"),
	}
}

func (r *PointsOfInterest) SearchPointsOfInterest(ctx context.Context, keyword string, limit int) ([]*ent.PointsOfInterest, error) {
	cli := r.GetTx(ctx).PointsOfInterest

	return cli.Query().
		Where(pointsofinterest.NameContains(keyword)).
		WithContinent().
		WithCountry().
		WithState().
		WithCity().
		Limit(limit).
		All(ctx)
}
