package impl

import (
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func NewCity(d *data.Data, logger *zap.SugaredLogger) repository.CityRepo {
	return &cityRepositoryImpl{
		Tx:                             d.Tx,
		logger:                         logger.Named("repo.city"),
		pointsOfInterestRepositoryImpl: new(pointsOfInterestRepositoryImpl),
		stateRepositoryImpl:            new(stateRepositoryImpl),
	}
}

type cityRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger

	pointsOfInterestRepositoryImpl repository.BaseRepo[*ent.PointsOfInterest, *do.PointsOfInterest]
	stateRepositoryImpl            repository.BaseRepo[*ent.State, *do.State]
}

func (c *cityRepositoryImpl) ToEntity(po *ent.City) *do.City {
	if po == nil {
		return nil
	}
	return &do.City{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		Name:      po.Name,
		NameEn:    po.NameEn,
		NameCn:    po.NameCn,
		Code:      po.Code,
		StateID:   po.StateID,
		Poi:       c.pointsOfInterestRepositoryImpl.ToEntities(po.Edges.Poi),
		State:     c.stateRepositoryImpl.ToEntity(po.Edges.State),
	}
}

func (c *cityRepositoryImpl) ToEntities(pos []*ent.City) []*do.City {
	if len(pos) == 0 {
		return nil
	}
	list := make([]*do.City, 0, len(pos))
	for _, v := range pos {
		list = append(list, c.ToEntity(v))
	}
	return list
}
