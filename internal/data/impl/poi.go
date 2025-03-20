package impl

import (
	"context"

	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/pointsofinterest"
	"github.com/iter-x/iter-x/internal/data/impl/build"
)

func NewPointsOfInterest(d *data.Data, cityRepository repository.CityRepo, logger *zap.SugaredLogger) repository.PointsOfInterestRepo {
	return &pointsOfInterestRepositoryImpl{
		Tx:             d.Tx,
		logger:         logger.Named("repo.poi"),
		cityRepository: cityRepository,
	}
}

type pointsOfInterestRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger

	cityRepository repository.CityRepo
}

func (r *pointsOfInterestRepositoryImpl) ToEntity(po *ent.PointsOfInterest) *do.PointsOfInterest {
	if po == nil {
		return nil
	}
	return build.PointsOfInterestRepositoryImplToEntity(po)
}

func (r *pointsOfInterestRepositoryImpl) ToEntities(pos []*ent.PointsOfInterest) []*do.PointsOfInterest {
	if len(pos) == 0 {
		return nil
	}
	return build.PointsOfInterestRepositoryImplToEntities(pos)
}

func (r *pointsOfInterestRepositoryImpl) SearchPointsOfInterest(ctx context.Context, params *bo.SearchPointsOfInterestParams) ([]*do.PointsOfInterest, error) {
	if !params.IsPoi() {
		return r.cityRepository.SearchPointsOfInterest(ctx, params)
	}
	cli := r.GetTx(ctx).PointsOfInterest
	keyword := params.Keyword
	limit := params.Limit
	rows, err := cli.Query().
		Where(pointsofinterest.Or(
			pointsofinterest.NameContains(keyword),
			pointsofinterest.NameCnContains(keyword),
			pointsofinterest.NameEnContains(keyword),
			pointsofinterest.DescriptionContains(keyword),
		)).
		WithContinent().
		WithCountry().
		WithState().
		WithCity().
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}
	pois := r.ToEntities(rows)
	otherRowLimit := limit - len(rows)
	if otherRowLimit > 0 && params.IsNext() {
		poiDos, err := r.cityRepository.SearchPointsOfInterest(ctx, params.DepthDec())
		if err != nil {
			return nil, err
		}
		pois = append(pois, poiDos...)
	}
	return pois, nil
}
