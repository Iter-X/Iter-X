package build

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func DailyTripRepositoryImplToEntity(po *ent.DailyTrip) *do.DailyTrip {
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

func DailyTripRepositoryImplToEntities(pos []*ent.DailyTrip) []*do.DailyTrip {
	if len(pos) == 0 {
		return nil
	}
	list := make([]*do.DailyTrip, 0, len(pos))
	for _, v := range pos {
		list = append(list, DailyTripRepositoryImplToEntity(v))
	}
	return list
}
