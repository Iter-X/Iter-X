package impl

import (
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func NewState(d *data.Data, logger *zap.SugaredLogger) repository.StateRepo {
	return &stateRepositoryImpl{
		Tx:                             d.Tx,
		logger:                         logger.Named("repo.state"),
		pointsOfInterestRepositoryImpl: new(pointsOfInterestRepositoryImpl),
		cityRepositoryImpl:             new(cityRepositoryImpl),
		countryRepositoryImpl:          new(countryRepositoryImpl),
	}
}

type stateRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger

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
