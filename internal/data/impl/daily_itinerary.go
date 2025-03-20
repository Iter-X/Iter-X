package impl

import (
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/impl/build"
)

func NewDailyItinerary(d *data.Data, logger *zap.SugaredLogger) repository.DailyItineraryRepo {
	return &dailyItineraryRepositoryImpl{
		Tx:     d.Tx,
		logger: logger.Named("repo.daily_itinerary"),
	}
}

type dailyItineraryRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger
}

func (d *dailyItineraryRepositoryImpl) ToEntity(po *ent.DailyItinerary) *do.DailyItinerary {
	if po == nil {
		return nil
	}
	return build.DailyItineraryRepositoryImplToEntity(po)
}

func (d *dailyItineraryRepositoryImpl) ToEntities(pos []*ent.DailyItinerary) []*do.DailyItinerary {
	if pos == nil {
		return nil
	}
	return build.DailyItineraryRepositoryImplToEntities(pos)
}
