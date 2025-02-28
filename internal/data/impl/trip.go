package impl

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/dailytrip"
	"github.com/iter-x/iter-x/internal/data/ent/trip"
)

type Trip struct {
	*data.Tx
	logger *zap.SugaredLogger
}

func NewTrip(tx *data.Tx, logger *zap.SugaredLogger) *Trip {
	return &Trip{
		Tx:     tx,
		logger: logger.Named("repo.trip"),
	}
}

func (r *Trip) CreateTrip(ctx context.Context, trip *ent.Trip) (*ent.Trip, error) {
	cli := r.GetTx(ctx).Trip

	return cli.Create().
		SetUserID(trip.UserID).
		SetTitle(trip.Title).
		SetDescription(trip.Description).
		SetStartDate(trip.StartDate).
		SetEndDate(trip.EndDate).
		Save(ctx)
}

func (r *Trip) GetTrip(ctx context.Context, id uuid.UUID) (*ent.Trip, error) {
	cli := r.GetTx(ctx).Trip

	return cli.Get(ctx, id)
}

func (r *Trip) UpdateTrip(ctx context.Context, trip *ent.Trip) (*ent.Trip, error) {
	cli := r.GetTx(ctx).Trip

	return cli.UpdateOneID(trip.ID).
		SetTitle(trip.Title).
		SetDescription(trip.Description).
		SetStartDate(trip.StartDate).
		SetEndDate(trip.EndDate).
		SetUpdatedAt(trip.UpdatedAt).
		Save(ctx)
}

func (r *Trip) DeleteTrip(ctx context.Context, id uuid.UUID) error {
	cli := r.GetTx(ctx).Trip

	return cli.DeleteOneID(id).Exec(ctx)
}

func (r *Trip) ListTrips(ctx context.Context, userId uuid.UUID) ([]*ent.Trip, error) {
	cli := r.GetTx(ctx).Trip

	return cli.Query().Where(trip.UserID(userId)).All(ctx)
}

func (r *Trip) CreateDailyTrip(ctx context.Context, dailyTrip *ent.DailyTrip) (*ent.DailyTrip, error) {
	cli := r.GetTx(ctx).DailyTrip

	return cli.Create().
		SetTripID(dailyTrip.TripID).
		SetDay(dailyTrip.Day).
		SetDate(dailyTrip.Date).
		SetNotes(dailyTrip.Notes).
		Save(ctx)
}

func (r *Trip) GetDailyTrip(ctx context.Context, tripId, dailyId uuid.UUID) (*ent.DailyTrip, error) {
	cli := r.GetTx(ctx).DailyTrip

	return cli.Query().
		Where(dailytrip.ID(dailyId), dailytrip.HasTripWith(trip.ID(tripId))).Only(ctx)
}

func (r *Trip) UpdateDailyTrip(ctx context.Context, dailyTrip *ent.DailyTrip) (*ent.DailyTrip, error) {
	cli := r.GetTx(ctx).DailyTrip

	return cli.UpdateOneID(dailyTrip.ID).
		SetDay(dailyTrip.Day).
		SetDate(dailyTrip.Date).
		SetNotes(dailyTrip.Notes).
		SetUpdatedAt(dailyTrip.UpdatedAt).
		Save(ctx)
}

func (r *Trip) DeleteDailyTrip(ctx context.Context, id uuid.UUID) error {
	cli := r.GetTx(ctx).DailyTrip

	return cli.DeleteOneID(id).Exec(ctx)
}

func (r *Trip) ListDailyTrips(ctx context.Context, tripId uuid.UUID) ([]*ent.DailyTrip, error) {
	cli := r.GetTx(ctx).DailyTrip

	return cli.Query().Where(dailytrip.TripID(tripId)).All(ctx)
}
