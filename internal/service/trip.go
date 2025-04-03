package service

import (
	"context"

	"github.com/google/uuid"

	tripV1 "github.com/iter-x/iter-x/internal/api/trip/v1"
	"github.com/iter-x/iter-x/internal/biz"
	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/common/cnst"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/service/build"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Trip struct {
	tripV1.UnimplementedTripServiceServer
	tripBiz *biz.Trip
}

func NewTrip(tripBiz *biz.Trip) *Trip {
	return &Trip{
		tripBiz: tripBiz,
	}
}

func validateTimeParams(duration int32, startTs, endTs *timestamppb.Timestamp) error {
	hasDuration := duration > 0
	hasTimeRange := startTs != nil && endTs != nil

	if !hasDuration && !hasTimeRange {
		return xerr.ErrorInvalidTimeRange()
	}
	if hasTimeRange && startTs.AsTime().After(endTs.AsTime()) {
		return xerr.ErrorInvalidTimeRange()
	}
	return nil
}

func (s *Trip) CreateTripManually(ctx context.Context, req *tripV1.CreateTripManuallyRequest) (*tripV1.CreateTripResponse, error) {
	if err := validateTimeParams(req.Duration, req.StartTs, req.EndTs); err != nil {
		return nil, err
	}

	trip, err := s.tripBiz.CreateTrip(ctx, &bo.CreateTripRequest{
		CreationMethod: cnst.TripCreationMethodManual,
		Destination:    req.Destination,
		StartDate:      req.StartTs.AsTime(),
		EndDate:        req.EndTs.AsTime(),
		Duration:       int(req.Duration),
	})
	if err != nil {
		return nil, err
	}
	return &tripV1.CreateTripResponse{Trip: build.ToTripProto(trip)}, nil
}

func (s *Trip) CreateTripFromCard(ctx context.Context, req *tripV1.CreateTripFromCardRequest) (*tripV1.CreateTripResponse, error) {
	if err := validateTimeParams(req.Duration, req.StartTs, req.EndTs); err != nil {
		return nil, err
	}
	if len(req.CityIds) == 0 {
		return nil, xerr.ErrorMissingCity()
	}

	trip, err := s.tripBiz.CreateTrip(ctx, &bo.CreateTripRequest{
		CreationMethod: cnst.TripCreationMethodCard,
		CityIds:        req.CityIds,
		PoiIds:         req.PoiIds,
		StartDate:      req.StartTs.AsTime(),
		EndDate:        req.EndTs.AsTime(),
		Duration:       int(req.Duration),
	})
	if err != nil {
		return nil, err
	}
	return &tripV1.CreateTripResponse{Trip: build.ToTripProto(trip)}, nil
}

func (s *Trip) CreateTripFromExternalLink(ctx context.Context, req *tripV1.CreateTripFromExternalLinkRequest) (*tripV1.CreateTripResponse, error) {
	if err := validateTimeParams(req.Duration, req.StartTs, req.EndTs); err != nil {
		return nil, err
	}
	if req.LinkUrl == "" {
		return nil, xerr.ErrorMissingLink()
	}

	trip, err := s.tripBiz.CreateTrip(ctx, &bo.CreateTripRequest{
		CreationMethod: cnst.TripCreationMethodExternalLink,
		LinkURL:        req.LinkUrl,
		StartDate:      req.StartTs.AsTime(),
		EndDate:        req.EndTs.AsTime(),
		Duration:       int(req.Duration),
	})
	if err != nil {
		return nil, err
	}
	return &tripV1.CreateTripResponse{Trip: build.ToTripProto(trip)}, nil
}

func (s *Trip) GetTrip(ctx context.Context, req *tripV1.GetTripRequest) (*tripV1.GetTripResponse, error) {
	tripId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}
	trip, err := s.tripBiz.GetTrip(ctx, tripId)
	if err != nil {
		return nil, err
	}
	return &tripV1.GetTripResponse{Trip: build.ToTripProto(trip)}, nil
}

func (s *Trip) UpdateTrip(ctx context.Context, req *tripV1.UpdateTripRequest) (*tripV1.UpdateTripResponse, error) {
	params := &bo.UpdateTripRequest{
		ID:          req.GetId(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		StartDate:   req.GetStartTs().AsTime(),
		EndDate:     req.GetEndTs().AsTime(),
		Status:      req.GetStatus(),
	}
	trip, err := s.tripBiz.UpdateTrip(ctx, params)
	if err != nil {
		return nil, err
	}
	return &tripV1.UpdateTripResponse{Trip: build.ToTripProto(trip)}, nil
}

func (s *Trip) DeleteTrip(ctx context.Context, req *tripV1.DeleteTripRequest) (*tripV1.DeleteTripResponse, error) {
	tripId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}
	err = s.tripBiz.DeleteTrip(ctx, tripId)
	if err != nil {
		return nil, err
	}
	return &tripV1.DeleteTripResponse{Status: "deleted"}, nil
}

func (s *Trip) ListTrips(ctx context.Context, _ *tripV1.ListTripsRequest) (*tripV1.ListTripsResponse, error) {
	trips, err := s.tripBiz.ListTrips(ctx)
	if err != nil {
		return nil, err
	}
	return &tripV1.ListTripsResponse{Trips: build.ToTripsProto(trips)}, nil
}

func (s *Trip) CreateDailyTrip(ctx context.Context, req *tripV1.CreateDailyTripRequest) (*tripV1.CreateDailyTripResponse, error) {
	params := &bo.CreateDailyTripRequest{
		TripID: req.GetTripId(),
		Date:   req.GetDate().AsTime(),
		Day:    req.GetDay(),
		Notes:  req.GetNotes(),
	}
	dailyTrip, err := s.tripBiz.CreateDailyTrip(ctx, params)
	if err != nil {
		return nil, err
	}
	return &tripV1.CreateDailyTripResponse{DailyTrip: build.ToDailyTripProto(dailyTrip)}, nil
}

func (s *Trip) GetDailyTrip(ctx context.Context, req *tripV1.GetDailyTripRequest) (*tripV1.GetDailyTripResponse, error) {
	params := &bo.GetDailyTripRequest{
		TripID:  req.GetTripId(),
		DailyID: req.GetDailyId(),
	}
	dailyTrip, err := s.tripBiz.GetDailyTrip(ctx, params)
	if err != nil {
		return nil, err
	}
	return &tripV1.GetDailyTripResponse{DailyTrip: build.ToDailyTripProto(dailyTrip)}, nil
}

func (s *Trip) UpdateDailyTrip(ctx context.Context, req *tripV1.UpdateDailyTripRequest) (*tripV1.UpdateDailyTripResponse, error) {
	params := &bo.UpdateDailyTripRequest{
		DailyID: req.GetDailyId(),
		Date:    req.GetDate().AsTime(),
		Day:     req.GetDay(),
		Notes:   req.GetNotes(),
		TripID:  req.GetTripId(),
	}
	dailyTrip, err := s.tripBiz.UpdateDailyTrip(ctx, params)
	if err != nil {
		return nil, err
	}
	return &tripV1.UpdateDailyTripResponse{DailyTrip: build.ToDailyTripProto(dailyTrip)}, nil
}

func (s *Trip) DeleteDailyTrip(ctx context.Context, req *tripV1.DeleteDailyTripRequest) (*tripV1.DeleteDailyTripResponse, error) {
	params := &bo.DeleteDailyTripRequest{
		DailyID: req.GetDailyId(),
		TripID:  req.GetTripId(),
	}
	if err := s.tripBiz.DeleteDailyTrip(ctx, params); err != nil {
		return nil, err
	}
	return &tripV1.DeleteDailyTripResponse{Status: "deleted"}, nil
}

func (s *Trip) ListDailyTrips(ctx context.Context, req *tripV1.ListDailyTripsRequest) (*tripV1.ListDailyTripsResponse, error) {
	params := &bo.ListDailyTripsRequest{
		TripID: req.GetTripId(),
	}
	dailyTrips, err := s.tripBiz.ListDailyTrips(ctx, params)
	if err != nil {
		return nil, err
	}
	return &tripV1.ListDailyTripsResponse{DailyTrips: build.ToDailyTripsProto(dailyTrips)}, nil
}
