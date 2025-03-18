package build

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	tripV1 "github.com/iter-x/iter-x/internal/api/trip/v1"
	"github.com/iter-x/iter-x/internal/biz/do"
)

func ToTripProto(t *do.Trip) *tripV1.Trip {
	return &tripV1.Trip{
		Id:        t.ID.String(),
		Status:    t.Status,
		Title:     t.Title,
		StartTs:   timestamppb.New(t.StartDate),
		EndTs:     timestamppb.New(t.EndDate),
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}

func ToTripsProto(ts []*do.Trip) []*tripV1.Trip {
	list := make([]*tripV1.Trip, 0, len(ts))
	for _, v := range ts {
		list = append(list, ToTripProto(v))
	}
	return list
}

func ToDailyTripProto(dt *do.DailyTrip) *tripV1.DailyTrip {
	return &tripV1.DailyTrip{
		Id:        dt.ID.String(),
		TripId:    dt.TripID.String(),
		Day:       dt.Day,
		Date:      timestamppb.New(dt.Date),
		Notes:     dt.Notes,
		CreatedAt: timestamppb.New(dt.CreatedAt),
		UpdatedAt: timestamppb.New(dt.UpdatedAt),
	}
}

func ToDailyTripsProto(dts []*do.DailyTrip) []*tripV1.DailyTrip {
	list := make([]*tripV1.DailyTrip, 0, len(dts))
	for _, v := range dts {
		list = append(list, ToDailyTripProto(v))
	}
	return list
}
