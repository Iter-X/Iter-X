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
		User:           AuthRepositoryImplToEntity(po.Edges.User),
		DailyTrip:      DailyTripRepositoryImplToEntities(po.Edges.DailyTrip),
		DailyItinerary: DailyItineraryRepositoryImplToEntities(po.Edges.DailyItinerary),
	}
}

func TripRepositoryImplToEntities(pos []*ent.Trip) []*do.Trip {
	if pos == nil {
		return nil
	}

	list := make([]*do.Trip, 0, len(pos))
	for _, v := range pos {
		list = append(list, TripRepositoryImplToEntity(v))
	}
	return list
}
