package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/api/trip/v1"
	"github.com/iter-x/iter-x/internal/biz"
	"github.com/iter-x/iter-x/internal/common/xerr"
)

type Trip struct {
	v1.UnimplementedTripServiceServer
	biz *biz.Trip
}

func NewTrip(biz *biz.Trip) *Trip {
	return &Trip{
		biz: biz,
	}
}

func (s *Trip) CreateTrip(ctx context.Context, req *v1.CreateTripRequest) (*v1.CreateTripResponse, error) {
	trip, err := s.biz.CreateTrip(ctx, req)
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
	trip, err := s.biz.GetTrip(ctx, tripId)
	if err != nil {
		return nil, err
	}
	return &v1.GetTripResponse{Trip: trip}, nil
}

func (s *Trip) UpdateTrip(ctx context.Context, req *v1.UpdateTripRequest) (*v1.UpdateTripResponse, error) {
	trip, err := s.biz.UpdateTrip(ctx, req)
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
	err = s.biz.DeleteTrip(ctx, tripId)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteTripResponse{Status: "deleted"}, nil
}

func (s *Trip) ListTrips(ctx context.Context, _ *v1.ListTripsRequest) (*v1.ListTripsResponse, error) {
	trips, err := s.biz.ListTrips(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ListTripsResponse{Trips: trips}, nil
}
