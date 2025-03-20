package impl

import (
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/impl/build"
)

func NewDailyTrip(d *data.Data, logger *zap.SugaredLogger) repository.DailyTripRepo {
	return &dailyTripRepositoryImpl{
		Tx:     d.Tx,
		logger: logger.Named("repo.daily_trip"),
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
	return build.DailyTripRepositoryImplToEntity(po)
}

func (d *dailyTripRepositoryImpl) ToEntities(pos []*ent.DailyTrip) []*do.DailyTrip {
	if len(pos) == 0 {
		return nil
	}
	return build.DailyTripRepositoryImplToEntities(pos)
}
