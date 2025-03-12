package impl

import (
	"context"

	"github.com/iter-x/iter-x/internal/data/ent/state"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func NewState(d *data.Data, countryRepository repository.CountryRepo, logger *zap.SugaredLogger) repository.StateRepo {
	return &stateRepositoryImpl{
		Tx:                             d.Tx,
		logger:                         logger.Named("repo.state"),
		countryRepository:              countryRepository,
		pointsOfInterestRepositoryImpl: new(pointsOfInterestRepositoryImpl),
		cityRepositoryImpl:             new(cityRepositoryImpl),
		countryRepositoryImpl:          new(countryRepositoryImpl),
	}
}

type stateRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger

	countryRepository repository.CountryRepo

	pointsOfInterestRepositoryImpl repository.BaseRepo[*ent.PointsOfInterest, *do.PointsOfInterest]
	cityRepositoryImpl             repository.BaseRepo[*ent.City, *do.City]
	countryRepositoryImpl          repository.BaseRepo[*ent.Country, *do.Country]
}

func (s *stateRepositoryImpl) ToEntity(po *ent.State) *do.State {
	if po == nil {
		return nil
	}
	return &do.State{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		Name:      po.Name,
		NameEn:    po.NameEn,
		NameCn:    po.NameCn,
		Code:      po.Code,
		CountryID: po.CountryID,
		Poi:       s.pointsOfInterestRepositoryImpl.ToEntities(po.Edges.Poi),
		City:      s.cityRepositoryImpl.ToEntities(po.Edges.City),
		Country:   s.countryRepositoryImpl.ToEntity(po.Edges.Country),
	}
}

func (s *stateRepositoryImpl) ToEntities(pos []*ent.State) []*do.State {
	if len(pos) == 0 {
		return nil
	}

	list := make([]*do.State, 0, len(pos))
	for _, v := range pos {
		list = append(list, s.ToEntity(v))
	}
	return list
}

func (s *stateRepositoryImpl) SearchPointsOfInterest(ctx context.Context, keyword string, limit int) ([]*do.PointsOfInterest, error) {
	cli := s.GetTx(ctx).State

	rows, err := cli.Query().
		Where(state.Or(
			state.NameContains(keyword),
			state.NameCnContains(keyword),
			state.NameEnContains(keyword),
		)).
		WithCountry(func(query *ent.CountryQuery) {
			query.WithContinent()
		}).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}
	pois := make([]*do.PointsOfInterest, 0, len(rows))
	for _, v := range rows {
		stateDo := s.ToEntity(v)
		pois = append(pois, &do.PointsOfInterest{
			State:     stateDo,
			Country:   stateDo.Country,
			Continent: stateDo.Country.Continent,
		})
	}
	otherRowLimit := limit - len(rows)
	if otherRowLimit > 0 {
		poiDos, err := s.countryRepository.SearchPointsOfInterest(ctx, keyword, otherRowLimit)
		if err != nil {
			return nil, err
		}
		pois = append(pois, poiDos...)
	}

	return pois, nil
}
