package impl

import (
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func NewCountry(tx *data.Tx, logger *zap.SugaredLogger) repository.CountryRepo {
	return &countryRepositoryImpl{
		Tx:                             tx,
		logger:                         logger,
		pointsOfInterestRepositoryImpl: new(pointsOfInterestRepositoryImpl),
	}
}

type countryRepositoryImpl struct {
	*data.Tx
	logger                         *zap.SugaredLogger
	pointsOfInterestRepositoryImpl repository.BaseRepo[*ent.PointsOfInterest, *do.PointsOfInterest]
}

func (c *countryRepositoryImpl) ToEntity(po *ent.Country) *do.Country {
	if po == nil {
		return nil
	}

	return &do.Country{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		Name:      po.Name,
		NameEn:    po.NameEn,
		NameCn:    po.NameCn,
		Poi:       c.pointsOfInterestRepositoryImpl.ToEntities(po.Edges.Poi),
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
