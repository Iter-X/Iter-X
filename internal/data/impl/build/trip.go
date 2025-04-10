package build

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func TripRepositoryImplToEntity(po *ent.Trip) *do.Trip {
	if po == nil {
		return nil
	}
	return &do.Trip{
		ID:             po.ID,
		CreatedAt:      po.CreatedAt,
		UpdatedAt:      po.UpdatedAt,
		UserID:         po.UserID,
		Status:         po.Status,
		Title:          po.Title,
		Description:    po.Description,
		StartDate:      po.StartDate,
		EndDate:        po.EndDate,
		Days:           po.Days,
		User:           AuthRepositoryImplToEntity(po.Edges.User),
		DailyTrip:      DailyTripRepositoryImplToEntities(po.Edges.DailyTrip),
		DailyItinerary: DailyItineraryRepositoryImplToEntities(po.Edges.DailyItinerary),
		PoiPool:        TripPOIPoolRepositoryImplToEntities(po.Edges.PoiPool),
	}
}

func TripRepositoryImplToEntities(pos []*ent.Trip) []*do.Trip {
	if pos == nil {
		return nil
	}
	list := make([]*do.Trip, 0, len(pos))
	for _, po := range pos {
		list = append(list, TripRepositoryImplToEntity(po))
	}
	return list
}

func DailyTripRepositoryImplToEntity(po *ent.DailyTrip) *do.DailyTrip {
	if po == nil {
		return nil
	}
	return &do.DailyTrip{
		ID:             po.ID,
		CreatedAt:      po.CreatedAt,
		UpdatedAt:      po.UpdatedAt,
		TripID:         po.TripID,
		Day:            po.Day,
		Date:           po.Date,
		Notes:          po.Notes,
		Trip:           TripRepositoryImplToEntity(po.Edges.Trip),
		DailyItinerary: DailyItineraryRepositoryImplToEntities(po.Edges.DailyItinerary),
	}
}

func DailyTripRepositoryImplToEntities(pos []*ent.DailyTrip) []*do.DailyTrip {
	if pos == nil {
		return nil
	}
	list := make([]*do.DailyTrip, 0, len(pos))
	for _, po := range pos {
		list = append(list, DailyTripRepositoryImplToEntity(po))
	}
	return list
}

func TripPOIPoolRepositoryImplToEntity(po *ent.TripPOIPool) *do.TripPOIPool {
	if po == nil {
		return nil
	}

	return &do.TripPOIPool{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		TripID:    po.TripID,
		PoiID:     po.PoiID,
		Trip:      TripRepositoryImplToEntity(po.Edges.Trip),
		Poi:       PointsOfInterestRepositoryImplToEntity(po.Edges.Poi),
	}
}

func TripPOIPoolRepositoryImplToEntities(pos []*ent.TripPOIPool) []*do.TripPOIPool {
	if len(pos) == 0 {
		return nil
	}

	entities := make([]*do.TripPOIPool, 0, len(pos))
	for _, po := range pos {
		entities = append(entities, TripPOIPoolRepositoryImplToEntity(po))
	}

	return entities
}
