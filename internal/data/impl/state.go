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
	}
}

type stateRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger

	pointsOfInterestRepositoryImpl repository.BaseRepo[*ent.PointsOfInterest, *do.PointsOfInterest]
}

func (s *stateRepositoryImpl) ToEntity(po *ent.State) *do.State {
	if po == nil {
		return nil
	}
	return &do.State{
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		ID:        po.ID,
		Name:      po.Name,
		NameCn:    po.NameCn,
		NameEn:    po.NameEn,
		Poi:       s.pointsOfInterestRepositoryImpl.ToEntities(po.Edges.Poi),
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
