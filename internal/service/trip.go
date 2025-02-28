package service

import (
	"context"

	"github.com/google/uuid"
	v1 "github.com/iter-x/iter-x/internal/api/trip/v1"
	"github.com/iter-x/iter-x/internal/biz"
	"github.com/iter-x/iter-x/internal/common/xerr"
)

type Trip struct {
	v1.UnimplementedTripServiceServer
	tripBiz *biz.Trip
}

func NewTrip(tripBiz *biz.Trip) *Trip {
	return &Trip{
		tripBiz: tripBiz,
	}
}

func (s *Trip) CreateTrip(ctx context.Context, req *v1.CreateTripRequest) (*v1.CreateTripResponse, error) {
	trip, err := s.tripBiz.CreateTrip(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.CreateTripResponse{Trip: trip}, nil
}

func (s *Trip) GetTrip(ctx context.Context, req *v1.GetTripRequest) (*v1.GetTripResponse, error) {
	tripId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}
	trip, err := s.tripBiz.GetTrip(ctx, tripId)
	if err != nil {
		return nil, err
	}
	return &v1.GetTripResponse{Trip: trip}, nil
}

func (s *Trip) UpdateTrip(ctx context.Context, req *v1.UpdateTripRequest) (*v1.UpdateTripResponse, error) {
	trip, err := s.tripBiz.UpdateTrip(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateTripResponse{Trip: trip}, nil
}

func (s *Trip) DeleteTrip(ctx context.Context, req *v1.DeleteTripRequest) (*v1.DeleteTripResponse, error) {
	tripId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}
	err = s.tripBiz.DeleteTrip(ctx, tripId)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteTripResponse{Status: "deleted"}, nil
}

func (s *Trip) ListTrips(ctx context.Context, _ *v1.ListTripsRequest) (*v1.ListTripsResponse, error) {
	trips, err := s.tripBiz.ListTrips(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ListTripsResponse{Trips: trips}, nil
}

func (s *Trip) CreateDailyTrip(ctx context.Context, req *v1.CreateDailyTripRequest) (*v1.CreateDailyTripResponse, error) {
	dailyTrip, err := s.tripBiz.CreateDailyTrip(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.CreateDailyTripResponse{DailyTrip: dailyTrip}, nil
}

func (s *Trip) GetDailyTrip(ctx context.Context, req *v1.GetDailyTripRequest) (*v1.GetDailyTripResponse, error) {
	dailyTrip, err := s.tripBiz.GetDailyTrip(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.GetDailyTripResponse{DailyTrip: dailyTrip}, nil
}

func (s *Trip) UpdateDailyTrip(ctx context.Context, req *v1.UpdateDailyTripRequest) (*v1.UpdateDailyTripResponse, error) {
	dailyTrip, err := s.tripBiz.UpdateDailyTrip(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateDailyTripResponse{DailyTrip: dailyTrip}, nil
}

func (s *Trip) DeleteDailyTrip(ctx context.Context, req *v1.DeleteDailyTripRequest) (*v1.DeleteDailyTripResponse, error) {
	err := s.tripBiz.DeleteDailyTrip(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteDailyTripResponse{Status: "deleted"}, nil
}

func (s *Trip) ListDailyTrips(ctx context.Context, req *v1.ListDailyTripsRequest) (*v1.ListDailyTripsResponse, error) {
	dailyTrips, err := s.tripBiz.ListDailyTrips(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.ListDailyTripsResponse{DailyTrips: dailyTrips}, nil
}
