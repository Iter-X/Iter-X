package impl

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/impl/ent"
	"go.uber.org/zap"
)

func NewTripDailyRepository(cli *ent.Client, logger *zap.SugaredLogger) repository.DailyTrip[*ent.DailyTrip, *do.DailyTrip] {
	return &tripDailyRepositoryImpl{
		Tx:                 &Tx{cli: cli},
		logger:             logger.Named("repo.DailyTrip"),
		authRepositoryImpl: new(authRepositoryImpl),
		tripRepositoryImpl: new(tripRepositoryImpl),
	}
}

type tripDailyRepositoryImpl struct {
	*Tx
	logger             *zap.SugaredLogger
	authRepositoryImpl repository.Base[*ent.User, *do.User]
	tripRepositoryImpl repository.Base[*ent.Trip, *do.Trip]
}

func (t *tripDailyRepositoryImpl) ToEntity(po *ent.DailyTrip) *do.DailyTrip {
	if po == nil {
		return nil
	}

	return &do.DailyTrip{
		ID:            po.ID,
		CreatedAt:     po.CreatedAt,
		UpdatedAt:     po.UpdatedAt,
		TripID:        po.TripID,
		Day:           po.Day,
		Date:          po.Date,
		Notes:         po.Notes,
		Trip:          t.tripRepositoryImpl.ToEntity(po.Edges.Trip),
		DailyTripItem: nil,
	}
}

func (t *tripDailyRepositoryImpl) ToEntities(pos []*ent.DailyTrip) []*do.DailyTrip {
	if len(pos) == 0 {
		return nil
	}
	list := make([]*do.DailyTrip, 0, len(pos))
	for _, v := range pos {
		list = append(list, t.ToEntity(v))
	}
	return list
}
