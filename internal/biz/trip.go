package biz

import (
	"context"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/helper/auth"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	"github.com/google/uuid"
	v1 "github.com/iter-x/iter-x/internal/api/trip/v1"
	"github.com/iter-x/iter-x/internal/repo"
	"github.com/iter-x/iter-x/internal/repo/ent"
	"go.uber.org/zap"
)

type Trip struct {
	repo   *repo.Trip
	logger *zap.SugaredLogger
}

func NewTrip(repo *repo.Trip, logger *zap.SugaredLogger) *Trip {
	return &Trip{
		repo:   repo,
		logger: logger.Named("biz.trip"),
	}
}

func (b *Trip) CreateTrip(ctx context.Context, req *v1.CreateTripRequest) (*v1.Trip, error) {
	claims, err := auth.ExtractClaims(ctx)
	if err != nil {
		return nil, xerr.ErrorUnauthorized()
	}

	trip := &ent.Trip{
		UserID:      claims.UID,
		Title:       req.Title,
		Description: req.Description,
		StartDate:   req.StartTs.AsTime(),
		EndDate:     req.EndTs.AsTime(),
	}

	createdTrip, err := b.repo.CreateTrip(ctx, trip)
	if err != nil {
		b.logger.Errorw("failed to create trip", "err", err)
		return nil, xerr.ErrorCreateTripFailed()
	}

	return toTripProto(createdTrip), nil
}

func (b *Trip) GetTrip(ctx context.Context, id uuid.UUID) (*v1.Trip, error) {
	trip, err := b.repo.GetTrip(ctx, id)
	if err != nil {
		b.logger.Errorw("failed to get trip", "err", err)
		return nil, xerr.ErrorGetTripFailed()
	}
	return toTripProto(trip), nil
}

func (b *Trip) UpdateTrip(ctx context.Context, req *v1.UpdateTripRequest) (*v1.Trip, error) {
	tripId, err := uuid.Parse(req.Id)
	if err != nil {
		b.logger.Errorw("failed to parse trip id", "err", err)
		return nil, xerr.ErrorInvalidTripId()
	}

	trip, err := b.repo.GetTrip(ctx, tripId)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, xerr.ErrorTripNotFound()
		}
		b.logger.Errorw("failed to get trip", "err", err)
		return nil, xerr.ErrorGetTripFailed()
	}

	trip.Title = req.Title
	trip.StartDate = req.StartTs.AsTime()
	trip.EndDate = req.EndTs.AsTime()
	trip.UpdatedAt = time.Now()

	updatedTrip, err := b.repo.UpdateTrip(ctx, trip)
	if err != nil {
		b.logger.Errorw("failed to update trip", "err", err)
		return nil, xerr.ErrorUpdateTripFailed()
	}

	return toTripProto(updatedTrip), nil
}

func (b *Trip) DeleteTrip(ctx context.Context, id uuid.UUID) error {
	if err := b.repo.DeleteTrip(ctx, id); err != nil {
		b.logger.Errorw("failed to delete trip", "err", err)
		return xerr.ErrorDeleteTripFailed()
	}
	return nil
}

func (b *Trip) ListTrips(ctx context.Context) ([]*v1.Trip, error) {
	claims, err := auth.ExtractClaims(ctx)
	if err != nil {
		return nil, xerr.ErrorUnauthorized()
	}

	trips, err := b.repo.ListTrips(ctx, claims.UID)
	if err != nil {
		b.logger.Errorw("failed to list trips", "err", err)
		return nil, xerr.ErrorGetTripListFailed()
	}

	tripList := make([]*v1.Trip, 0, len(trips))
	for _, t := range trips {
		tripList = append(tripList, toTripProto(t))
	}

	return tripList, nil
}

func (b *Trip) CreateDailyTrip(ctx context.Context, req *v1.CreateDailyTripRequest) (*v1.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripId)
	if err != nil {
		b.logger.Errorw("failed to parse trip id", "err", err)
		return nil, xerr.ErrorInvalidTripId()
	}

	_, err = b.repo.GetTrip(ctx, tripId)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, xerr.ErrorTripNotFound()
		}
		b.logger.Errorw("failed to get trip", "err", err)
		return nil, xerr.ErrorGetTripFailed()
	}

	dailyTrip := &ent.DailyTrip{
		TripID: tripId,
		Day:    req.Day,
		Date:   req.Date.AsTime(),
		Notes:  req.Notes,
	}

	createdDailyTrip, err := b.repo.CreateDailyTrip(ctx, dailyTrip)
	if err != nil {
		b.logger.Errorw("failed to create daily trip", "err", err)
		return nil, xerr.ErrorCreateDailyTripFailed()
	}

	return toDailyTripProto(createdDailyTrip), nil
}

func (b *Trip) GetDailyTrip(ctx context.Context, req *v1.GetDailyTripRequest) (*v1.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripId)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}
	dailyId, err := uuid.Parse(req.DailyId)
	if err != nil {
		return nil, xerr.ErrorInvalidDailyTripId()
	}

	dailyTrip, err := b.repo.GetDailyTrip(ctx, tripId, dailyId)
	if err != nil {
		b.logger.Errorw("failed to get daily trip", "err", err)
		return nil, xerr.ErrorGetDailyTripFailed()
	}
	return toDailyTripProto(dailyTrip), nil
}

func (b *Trip) UpdateDailyTrip(ctx context.Context, req *v1.UpdateDailyTripRequest) (*v1.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripId)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}
	dailyId, err := uuid.Parse(req.DailyId)
	if err != nil {
		return nil, xerr.ErrorInvalidDailyTripId()
	}

	dailyTrip, err := b.repo.GetDailyTrip(ctx, tripId, dailyId)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, xerr.ErrorDailyTripNotFound()
		}
		b.logger.Errorw("daily trip not found", "err", err)
		return nil, xerr.ErrorGetDailyTripFailed()
	}

	dailyTrip.Day = req.Day
	dailyTrip.Date = req.Date.AsTime()
	dailyTrip.Notes = req.Notes
	dailyTrip.UpdatedAt = time.Now()

	updatedDailyTrip, err := b.repo.UpdateDailyTrip(ctx, dailyTrip)
	if err != nil {
		b.logger.Errorw("failed to update daily trip", "err", err)
		return nil, xerr.ErrorUpdateDailyTripFailed()
	}

	return toDailyTripProto(updatedDailyTrip), nil
}

func (b *Trip) DeleteDailyTrip(ctx context.Context, req *v1.DeleteDailyTripRequest) error {
	dailyId, err := uuid.Parse(req.DailyId)
	if err != nil {
		return xerr.ErrorInvalidDailyTripId()
	}

	if err := b.repo.DeleteDailyTrip(ctx, dailyId); err != nil {
		b.logger.Errorw("failed to delete daily trip", "err", err)
		return xerr.ErrorDeleteDailyTripFailed()
	}
	return nil
}

func (b *Trip) ListDailyTrips(ctx context.Context, req *v1.ListDailyTripsRequest) ([]*v1.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripId)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}

	dailyTrips, err := b.repo.ListDailyTrips(ctx, tripId)
	if err != nil {
		b.logger.Errorw("failed to list daily trips", "err", err)
		return nil, xerr.ErrorGetDailyTripListFailed()
	}

	dailyTripList := make([]*v1.DailyTrip, 0, len(dailyTrips))
	for _, dt := range dailyTrips {
		dailyTripList = append(dailyTripList, toDailyTripProto(dt))
	}

	return dailyTripList, nil
}

func toDailyTripProto(dt *ent.DailyTrip) *v1.DailyTrip {
	return &v1.DailyTrip{
		Id:        dt.ID.String(),
		TripId:    dt.TripID.String(),
		Day:       dt.Day,
		Date:      timestamppb.New(dt.Date),
		Notes:     dt.Notes,
		CreatedAt: timestamppb.New(dt.CreatedAt),
		UpdatedAt: timestamppb.New(dt.UpdatedAt),
	}
}

func toTripProto(t *ent.Trip) *v1.Trip {
	return &v1.Trip{
		Id:        t.ID.String(),
		Status:    t.Status,
		Title:     t.Title,
		StartTs:   timestamppb.New(t.StartDate),
		EndTs:     timestamppb.New(t.EndDate),
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}
