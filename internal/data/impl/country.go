package impl

import (
	"context"

	"github.com/iter-x/iter-x/internal/data/ent/country"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func NewCountry(d *data.Data, continentRepository repository.ContinentRepo, logger *zap.SugaredLogger) repository.CountryRepo {
	return &countryRepositoryImpl{
		Tx:                             d.Tx,
		logger:                         logger.Named("repo.country"),
		continentRepository:            continentRepository,
		pointsOfInterestRepositoryImpl: new(pointsOfInterestRepositoryImpl),
		stateRepositoryImpl:            new(stateRepositoryImpl),
		continentRepositoryImpl:        new(continentRepositoryImpl),
	}
}

type countryRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger

	continentRepository repository.ContinentRepo

	pointsOfInterestRepositoryImpl repository.BaseRepo[*ent.PointsOfInterest, *do.PointsOfInterest]
	stateRepositoryImpl            repository.BaseRepo[*ent.State, *do.State]
	continentRepositoryImpl        repository.BaseRepo[*ent.Continent, *do.Continent]
}

func (c *countryRepositoryImpl) ToEntity(po *ent.Country) *do.Country {
	if po == nil {
		return nil
	}

	return &do.Country{
		ID:          po.ID,
		CreatedAt:   po.CreatedAt,
		UpdatedAt:   po.UpdatedAt,
		Name:        po.Name,
		NameEn:      po.NameEn,
		NameCn:      po.NameCn,
		Code:        po.Code,
		ContinentID: po.ContinentID,
		Poi:         c.pointsOfInterestRepositoryImpl.ToEntities(po.Edges.Poi),
		State:       c.stateRepositoryImpl.ToEntities(po.Edges.State),
		Continent:   c.continentRepositoryImpl.ToEntity(po.Edges.Continent),
	}
}

func (c *countryRepositoryImpl) ToEntities(pos []*ent.Country) []*do.Country {
	if len(pos) == 0 {
		return nil
	}
	list := make([]*do.Country, 0, len(pos))
	for _, v := range pos {
		list = append(list, c.ToEntity(v))
	}
	return list
}

func (c *countryRepositoryImpl) SearchPointsOfInterest(ctx context.Context, keyword string, limit int) ([]*do.PointsOfInterest, error) {
	cli := c.GetTx(ctx).Country

	rows, err := cli.Query().
		Where(country.Or(
			country.NameContains(keyword),
			country.NameCnContains(keyword),
			country.NameEnContains(keyword),
			country.CodeContains(keyword),
		)).
		WithPoi().
		WithState().
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return c.continentRepository.SearchPointsOfInterest(ctx, keyword, limit)
	}

	pois := make([]*do.PointsOfInterest, 0, len(rows))
	for _, v := range rows {
		countryDo := c.ToEntity(v)
		pois = append(pois, &do.PointsOfInterest{
			Country:   countryDo,
			Continent: countryDo.Continent,
		})
	}
	return pois, nil
}
