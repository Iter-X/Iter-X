package biz

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/ai/agent"
	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/common/cnst"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/helper/auth"
)

type Trip struct {
	tripRepo repository.TripRepo
	logger   *zap.SugaredLogger
	agentHub *agent.Hub
}

func NewTrip(tripRepo repository.TripRepo, logger *zap.SugaredLogger, agentHub *agent.Hub) *Trip {
	return &Trip{
		tripRepo: tripRepo,
		logger:   logger.Named("biz.tripRepo"),
		agentHub: agentHub,
	}
}

func (b *Trip) CreateTrip(ctx context.Context, req *bo.CreateTripRequest) (*do.Trip, error) {
	claims, err := auth.ExtractClaims(ctx)
	if err != nil {
		return nil, xerr.ErrorUnauthorized()
	}

	switch req.CreationMethod {
	case cnst.TripCreationMethodManual:
		// Get the PlanAgent from the agentHub
		planAgent, err := b.agentHub.GetAgent("PlanAgent")
		if err != nil {
			b.logger.Errorw("failed to get PlanAgent", "err", err)
			return nil, xerr.ErrorCreateTripFailed()
		}

		// Execute the PlanAgent
		result, err := planAgent.Execute(ctx, &do.PlanAgentInput{
			Destination: req.Destination,
			StartDate:   req.StartDate,
			EndDate:     req.EndDate,
			Duration:    req.Duration,
		})
		if err != nil {
			b.logger.Errorw("failed to plan trip with PlanAgent", "err", err)
			return nil, xerr.ErrorCreateTripFailed()
		}
		if result == nil {
			b.logger.Error("failed to plan trip with PlanAgent because result is nil")
			return nil, xerr.ErrorCreateTripFailed()
		}
		rawTrip, ok := result.(*do.PlanAgentOutput)
		if !ok {
			b.logger.Error("failed to cast PlanAgent output to PlanAgentOutput")
			return nil, xerr.ErrorCreateTripFailed()
		}

		// Create the trip
		_ = claims.UID
		_ = rawTrip
		return nil, nil
	case cnst.TripCreationMethodCard:
		// TODO: Handle card-based creation
		return nil, xerr.ErrorCreateTripFailed()
	case cnst.TripCreationMethodExternalLink:
		// TODO: Handle external link creation
		return nil, xerr.ErrorCreateTripFailed()
	case cnst.TripCreationMethodImage:
		// TODO: Handle image-based creation
		return nil, xerr.ErrorCreateTripFailed()
	case cnst.TripCreationMethodVoice:
		// TODO: Handle voice-based creation
		return nil, xerr.ErrorCreateTripFailed()
	default:
		return nil, xerr.ErrorInvalidCreationMethod()
	}
}

func (b *Trip) GetTrip(ctx context.Context, id uuid.UUID) (*do.Trip, error) {
	trip, err := b.tripRepo.GetTrip(ctx, id)
	if err != nil {
		b.logger.Errorw("failed to get tripRepo", "err", err)
		return nil, xerr.ErrorGetTripFailed()
	}
	return trip, nil
}

func (b *Trip) UpdateTrip(ctx context.Context, req *bo.UpdateTripRequest) (*do.Trip, error) {
	tripId, err := uuid.Parse(req.ID)
	if err != nil {
		b.logger.Errorw("failed to parse tripRepo id", "err", err)
		return nil, xerr.ErrorInvalidTripId()
	}

	trip, err := b.tripRepo.GetTrip(ctx, tripId)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, xerr.ErrorTripNotFound()
		}
		b.logger.Errorw("failed to get tripRepo", "err", err)
		return nil, xerr.ErrorGetTripFailed()
	}

	trip.Title = req.Title
	trip.StartDate = req.StartDate
	trip.EndDate = req.EndDate
	trip.UpdatedAt = time.Now()

	updatedTrip, err := b.tripRepo.UpdateTrip(ctx, trip)
	if err != nil {
		b.logger.Errorw("failed to update tripRepo", "err", err)
		return nil, xerr.ErrorUpdateTripFailed()
	}

	return updatedTrip, nil
}

func (b *Trip) DeleteTrip(ctx context.Context, id uuid.UUID) error {
	if err := b.tripRepo.DeleteTrip(ctx, id); err != nil {
		b.logger.Errorw("failed to delete tripRepo", "err", err)
		return xerr.ErrorDeleteTripFailed()
	}
	return nil
}

func (b *Trip) ListTrips(ctx context.Context) ([]*do.Trip, error) {
	claims, err := auth.ExtractClaims(ctx)
	if err != nil {
		return nil, xerr.ErrorUnauthorized()
	}

	trips, err := b.tripRepo.ListTrips(ctx, claims.UID)
	if err != nil {
		b.logger.Errorw("failed to list trips", "err", err)
		return nil, xerr.ErrorGetTripListFailed()
	}
	return trips, nil
}

func (b *Trip) CreateDailyTrip(ctx context.Context, req *bo.CreateDailyTripRequest) (*do.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripID)
	if err != nil {
		b.logger.Errorw("failed to parse tripRepo id", "err", err)
		return nil, xerr.ErrorInvalidTripId()
	}

	_, err = b.tripRepo.GetTrip(ctx, tripId)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, xerr.ErrorTripNotFound()
		}
		b.logger.Errorw("failed to get tripRepo", "err", err)
		return nil, xerr.ErrorGetTripFailed()
	}

	dailyTrip := &do.DailyTrip{
		TripID: tripId,
		Day:    req.Day,
		Date:   req.Date,
		Notes:  req.Notes,
	}

	createdDailyTrip, err := b.tripRepo.CreateDailyTrip(ctx, dailyTrip)
	if err != nil {
		b.logger.Errorw("failed to create daily tripRepo", "err", err)
		return nil, xerr.ErrorCreateDailyTripFailed()
	}

	return createdDailyTrip, nil
}

func (b *Trip) GetDailyTrip(ctx context.Context, req *bo.GetDailyTripRequest) (*do.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripID)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}
	dailyId, err := uuid.Parse(req.DailyID)
	if err != nil {
		return nil, xerr.ErrorInvalidDailyTripId()
	}

	dailyTrip, err := b.tripRepo.GetDailyTrip(ctx, tripId, dailyId)
	if err != nil {
		b.logger.Errorw("failed to get daily tripRepo", "err", err)
		return nil, xerr.ErrorGetDailyTripFailed()
	}
	return dailyTrip, nil
}

func (b *Trip) UpdateDailyTrip(ctx context.Context, req *bo.UpdateDailyTripRequest) (*do.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripID)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}
	dailyId, err := uuid.Parse(req.DailyID)
	if err != nil {
		return nil, xerr.ErrorInvalidDailyTripId()
	}

	dailyTrip, err := b.tripRepo.GetDailyTrip(ctx, tripId, dailyId)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, xerr.ErrorDailyTripNotFound()
		}
		b.logger.Errorw("daily tripRepo not found", "err", err)
		return nil, xerr.ErrorGetDailyTripFailed()
	}

	dailyTrip.Day = req.Day
	dailyTrip.Date = req.Date
	dailyTrip.Notes = req.Notes
	dailyTrip.UpdatedAt = time.Now()

	updatedDailyTrip, err := b.tripRepo.UpdateDailyTrip(ctx, dailyTrip)
	if err != nil {
		b.logger.Errorw("failed to update daily tripRepo", "err", err)
		return nil, xerr.ErrorUpdateDailyTripFailed()
	}

	return updatedDailyTrip, nil
}

func (b *Trip) DeleteDailyTrip(ctx context.Context, req *bo.DeleteDailyTripRequest) error {
	dailyId, err := uuid.Parse(req.DailyID)
	if err != nil {
		return xerr.ErrorInvalidDailyTripId()
	}

	if err := b.tripRepo.DeleteDailyTrip(ctx, dailyId); err != nil {
		b.logger.Errorw("failed to delete daily tripRepo", "err", err)
		return xerr.ErrorDeleteDailyTripFailed()
	}
	return nil
}

func (b *Trip) ListDailyTrips(ctx context.Context, req *bo.ListDailyTripsRequest) ([]*do.DailyTrip, error) {
	tripId, err := uuid.Parse(req.TripID)
	if err != nil {
		return nil, xerr.ErrorInvalidTripId()
	}

	dailyTrips, err := b.tripRepo.ListDailyTrips(ctx, tripId)
	if err != nil {
		b.logger.Errorw("failed to list daily trips", "err", err)
		return nil, xerr.ErrorGetDailyTripListFailed()
	}

	return dailyTrips, nil
}
