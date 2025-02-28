package impl

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/impl/ent"
	"go.uber.org/zap"
)

func NewTripDailyItemRepository(cli *ent.Client, logger *zap.SugaredLogger) repository.DailyTripItem[*ent.DailyTripItem, *do.DailyTripItem] {
	return &tripDailyItemRepositoryImpl{
		Tx:                      &Tx{cli: cli},
		logger:                  logger,
		tripRepositoryImpl:      new(tripRepositoryImpl),
		tripDailyRepositoryImpl: new(tripDailyRepositoryImpl),
	}
}

type tripDailyItemRepositoryImpl struct {
	*Tx
	logger *zap.SugaredLogger

	tripRepositoryImpl      repository.Base[*ent.Trip, *do.Trip]
	tripDailyRepositoryImpl repository.Base[*ent.DailyTrip, *do.DailyTrip]
}

func (t *tripDailyItemRepositoryImpl) ToEntity(po *ent.DailyTripItem) *do.DailyTripItem {
	if po == nil {
		return nil
	}
	return &do.DailyTripItem{
		ID:          po.ID,
		CreatedAt:   po.CreatedAt,
		UpdatedAt:   po.UpdatedAt,
		TripID:      po.TripID,
		DailyTripID: po.DailyTripID,
		Notes:       po.Notes,
		Trip:        t.tripRepositoryImpl.ToEntity(po.Edges.Trip),
		DailyTrip:   t.tripDailyRepositoryImpl.ToEntity(po.Edges.DailyTrip),
	}
}

func (t *tripDailyItemRepositoryImpl) ToEntities(pos []*ent.DailyTripItem) []*do.DailyTripItem {
	if len(pos) == 0 {
		return nil
	}

	list := make([]*do.DailyTripItem, 0, len(pos))
	for _, v := range pos {
		list = append(list, t.ToEntity(v))
	}
	return list
}
