package impl

import (
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func NewDailyItinerary(d *data.Data, logger *zap.SugaredLogger) repository.DailyItineraryRepo {
	return &dailyItineraryRepositoryImpl{
		Tx:                             d.Tx,
		logger:                         logger.Named("repo.daily_itinerary"),
		dailyTripRepositoryImpl:        new(dailyTripRepositoryImpl),
		pointsOfInterestRepositoryImpl: new(pointsOfInterestRepositoryImpl),
		tripRepositoryImpl:             new(tripRepositoryImpl),
	}
}

type dailyItineraryRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger

	tripRepositoryImpl             repository.BaseRepo[*ent.Trip, *do.Trip]
	pointsOfInterestRepositoryImpl repository.BaseRepo[*ent.PointsOfInterest, *do.PointsOfInterest]
	dailyTripRepositoryImpl        repository.BaseRepo[*ent.DailyTrip, *do.DailyTrip]
}

func (d *dailyItineraryRepositoryImpl) ToEntity(po *ent.DailyItinerary) *do.DailyItinerary {
	if po == nil {
		return nil
	}
	return &do.DailyItinerary{
		ID:          po.ID,
		CreatedAt:   po.CreatedAt,
		UpdatedAt:   po.UpdatedAt,
		TripID:      po.TripID,
		DailyTripID: po.DailyTripID,
		PoiID:       po.PoiID,
		Notes:       po.Notes,
		Trip:        d.tripRepositoryImpl.ToEntity(po.Edges.Trip),
		DailyTrip:   d.dailyTripRepositoryImpl.ToEntity(po.Edges.DailyTrip),
		Poi:         d.pointsOfInterestRepositoryImpl.ToEntity(po.Edges.Poi),
	}
}

func (d *dailyItineraryRepositoryImpl) ToEntities(pos []*ent.DailyItinerary) []*do.DailyItinerary {
	if pos == nil {
		return nil
	}
	list := make([]*do.DailyItinerary, 0, len(pos))
	for _, v := range pos {
		list = append(list, d.ToEntity(v))
	}
	return list
}
