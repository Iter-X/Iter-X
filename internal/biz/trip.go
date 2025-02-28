package biz

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	tripV1 "github.com/iter-x/iter-x/internal/api/trip/v1"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/helper/auth"
	"github.com/iter-x/iter-x/internal/impl/ent"
)

type Trip struct {
	tripRepo repository.TripRepo
	logger   *zap.SugaredLogger
}

func NewTrip(tripRepo repository.TripRepo, logger *zap.SugaredLogger) *Trip {
	return &Trip{
		tripRepo: tripRepo,
		logger:   logger.Named("biz.trip"),
	}
}

func (b *Trip) CreateTrip(ctx context.Context, req *tripV1.CreateTripRequest) (*tripV1.Trip, error) {
	claims, err := auth.ExtractClaims(ctx)
	if err != nil {
		return nil, xerr.ErrorUnauthorized()
	}

	trip := &do.Trip{
		UserID:      claims.UID,
		Title:       req.Title,
		Description: req.Description,
		StartDate:   req.StartTs.AsTime(),
		EndDate:     req.EndTs.AsTime(),
	}

	createdTrip, err := b.tripRepo.CreateTrip(ctx, trip)
	if err != nil {
		b.logger.Errorw("failed to create trip", "err", err)
		return nil, xerr.ErrorCreateTripFailed()
	}

	return createdTrip.ToTripProto(), nil
}

func (b *Trip) GetTrip(ctx context.Context, id uuid.UUID) (*tripV1.Trip, error) {
	trip, err := b.tripRepo.GetTrip(ctx, id)
	if err != nil {
		b.logger.Errorw("failed to get trip", "err", err)
		return nil, xerr.ErrorGetTripFailed()
	}
	return trip.ToTripProto(), nil
}

func (b *Trip) UpdateTrip(ctx context.Context, req *tripV1.UpdateTripRequest) (*tripV1.Trip, error) {
	tripId, err := uuid.Parse(req.Id)
	if err != nil {
		b.logger.Errorw("failed to parse trip id", "err", err)
		return nil, xerr.ErrorInvalidTripId()
	}

	trip, err := b.tripRepo.GetTrip(ctx, tripId)
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

	updatedTrip, err := b.tripRepo.UpdateTrip(ctx, trip)
	if err != nil {
		b.logger.Errorw("failed to update trip", "err", err)
		return nil, xerr.ErrorUpdateTripFailed()
	}

	return updatedTrip.ToTripProto(), nil
}

func (b *Trip) DeleteTrip(ctx context.Context, id uuid.UUID) error {
	if err := b.tripRepo.DeleteTrip(ctx, id); err != nil {
		b.logger.Errorw("failed to delete trip", "err", err)
		return xerr.ErrorDeleteTripFailed()
	}
	return nil
}

func (b *Trip) ListTrips(ctx context.Context) ([]*tripV1.Trip, error) {
	claims, err := auth.ExtractClaims(ctx)
	if err != nil {
		return nil, xerr.ErrorUnauthorized()
	}

	trips, err := b.tripRepo.ListTrips(ctx, claims.UID)
	if err != nil {
		b.logger.Errorw("failed to list trips", "err", err)
		return nil, xerr.ErrorGetTripListFailed()
	}

	tripList := make([]*tripV1.Trip, 0, len(trips))
	for _, t := range trips {
		tripList = append(tripList, t.ToTripProto())
	}

	return tripList, nil
}

func (b *Trip) CreateDailyTrip(ctx context.Context, req *tripV1.CreateDailyTripRequest) (*tripV1.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripId)
	if err != nil {
		b.logger.Errorw("failed to parse trip id", "err", err)
		return nil, xerr.ErrorInvalidTripId()
	}

	_, err = b.tripRepo.GetTrip(ctx, tripId)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, xerr.ErrorTripNotFound()
		}
		b.logger.Errorw("failed to get trip", "err", err)
		return nil, xerr.ErrorGetTripFailed()
	}

	dailyTrip := &do.DailyTrip{
		TripID: tripId,
		Day:    req.Day,
		Date:   req.Date.AsTime(),
		Notes:  req.Notes,
	}

	createdDailyTrip, err := b.tripRepo.CreateDailyTrip(ctx, dailyTrip)
	if err != nil {
		b.logger.Errorw("failed to create daily trip", "err", err)
		return nil, xerr.ErrorCreateDailyTripFailed()
	}

	return createdDailyTrip.ToDailyTripProto(), nil
}

func (b *Trip) GetDailyTrip(ctx context.Context, req *tripV1.GetDailyTripRequest) (*tripV1.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripId)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}
	dailyId, err := uuid.Parse(req.DailyId)
	if err != nil {
		return nil, xerr.ErrorInvalidDailyTripId()
	}

	dailyTrip, err := b.tripRepo.GetDailyTrip(ctx, tripId, dailyId)
	if err != nil {
		b.logger.Errorw("failed to get daily trip", "err", err)
		return nil, xerr.ErrorGetDailyTripFailed()
	}
	return dailyTrip.ToDailyTripProto(), nil
}

func (b *Trip) UpdateDailyTrip(ctx context.Context, req *tripV1.UpdateDailyTripRequest) (*tripV1.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripId)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}
	dailyId, err := uuid.Parse(req.DailyId)
	if err != nil {
		return nil, xerr.ErrorInvalidDailyTripId()
	}

	dailyTrip, err := b.tripRepo.GetDailyTrip(ctx, tripId, dailyId)
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

	updatedDailyTrip, err := b.tripRepo.UpdateDailyTrip(ctx, dailyTrip)
	if err != nil {
		b.logger.Errorw("failed to update daily trip", "err", err)
		return nil, xerr.ErrorUpdateDailyTripFailed()
	}

	return updatedDailyTrip.ToDailyTripProto(), nil
}

func (b *Trip) DeleteDailyTrip(ctx context.Context, req *tripV1.DeleteDailyTripRequest) error {
	dailyId, err := uuid.Parse(req.DailyId)
	if err != nil {
		return xerr.ErrorInvalidDailyTripId()
	}

	if err := b.tripRepo.DeleteDailyTrip(ctx, dailyId); err != nil {
		b.logger.Errorw("failed to delete daily trip", "err", err)
		return xerr.ErrorDeleteDailyTripFailed()
	}
	return nil
}

func (b *Trip) ListDailyTrips(ctx context.Context, req *tripV1.ListDailyTripsRequest) ([]*tripV1.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripId)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}

	dailyTrips, err := b.tripRepo.ListDailyTrips(ctx, tripId)
	if err != nil {
		b.logger.Errorw("failed to list daily trips", "err", err)
		return nil, xerr.ErrorGetDailyTripListFailed()
	}

	dailyTripList := make([]*tripV1.DailyTrip, 0, len(dailyTrips))
	for _, dt := range dailyTrips {
		dailyTripList = append(dailyTripList, dt.ToDailyTripProto())
	}

	return dailyTripList, nil
}
