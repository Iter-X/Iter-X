package build

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func DailyItineraryRepositoryImplToEntity(po *ent.DailyItinerary) *do.DailyItinerary {
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
		Trip:        TripRepositoryImplToEntity(po.Edges.Trip),
		DailyTrip:   DailyTripRepositoryImplToEntity(po.Edges.DailyTrip),
		Poi:         PointsOfInterestRepositoryImplToEntity(po.Edges.Poi),
	}
}

func DailyItineraryRepositoryImplToEntities(pos []*ent.DailyItinerary) []*do.DailyItinerary {
	if pos == nil {
		return nil
	}
	list := make([]*do.DailyItinerary, 0, len(pos))
	for _, v := range pos {
		list = append(list, DailyItineraryRepositoryImplToEntity(v))
	}
	return list
}
