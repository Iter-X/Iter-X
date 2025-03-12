package impl

import (
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func NewContinent(d *data.Data, logger *zap.SugaredLogger) repository.ContinentRepo {
	return &continentRepositoryImpl{
		Tx:                             d.Tx,
		logger:                         logger.Named("repo.continent"),
		pointsOfInterestRepositoryImpl: new(pointsOfInterestRepositoryImpl),
		countryRepositoryImpl:          new(countryRepositoryImpl),
	}
}

type continentRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger

	pointsOfInterestRepositoryImpl repository.BaseRepo[*ent.PointsOfInterest, *do.PointsOfInterest]
	countryRepositoryImpl          repository.BaseRepo[*ent.Country, *do.Country]
}

func (c *continentRepositoryImpl) ToEntity(po *ent.Continent) *do.Continent {
	if po == nil {
		return nil
	}
	return &do.Continent{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		Name:      po.Name,
		NameEn:    po.NameEn,
		NameCn:    po.NameCn,
		Code:      po.Code,
		Poi:       c.pointsOfInterestRepositoryImpl.ToEntities(po.Edges.Poi),
		Country:   c.countryRepositoryImpl.ToEntities(po.Edges.Country),
	}
}

func (c *continentRepositoryImpl) ToEntities(pos []*ent.Continent) []*do.Continent {
	if pos == nil {
		return nil
	}
	list := make([]*do.Continent, 0, len(pos))
	for _, v := range pos {
		list = append(list, c.ToEntity(v))
	}
	return list
}
