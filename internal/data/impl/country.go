package impl

import (
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func NewCountry(d *data.Data, logger *zap.SugaredLogger) repository.CountryRepo {
	return &countryRepositoryImpl{
		Tx:                             d.Tx,
		logger:                         logger.Named("repo.country"),
		pointsOfInterestRepositoryImpl: new(pointsOfInterestRepositoryImpl),
		stateRepositoryImpl:            new(stateRepositoryImpl),
		continentRepositoryImpl:        new(continentRepositoryImpl),
	}
}

type countryRepositoryImpl struct {
	*data.Tx
	logger                         *zap.SugaredLogger
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
