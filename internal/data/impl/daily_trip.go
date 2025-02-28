package impl

import (
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func NewDailyTrip(tx *data.Tx, logger *zap.SugaredLogger) repository.DailyTripRepo {
	return &dailyTripRepositoryImpl{
		Tx:     tx,
		logger: logger,
	}
}

type dailyTripRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger
}

func (d *dailyTripRepositoryImpl) ToEntity(po *ent.DailyTrip) *do.DailyTrip {
	if po == nil {
		return nil
	}
	return &do.DailyTrip{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		TripID:    po.TripID,
		Day:       po.Day,
		Date:      po.Date,
		Notes:     po.Notes,
	}
}

func (d *dailyTripRepositoryImpl) ToEntities(pos []*ent.DailyTrip) []*do.DailyTrip {
	if len(pos) == 0 {
		return nil
	}
	list := make([]*do.DailyTrip, 0, len(pos))
	for _, v := range pos {
		list = append(list, d.ToEntity(v))
	}
	return list
}
