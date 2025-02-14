package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/repo/ent"
	"github.com/iter-x/iter-x/internal/repo/ent/dailytrip"
	"github.com/iter-x/iter-x/internal/repo/ent/trip"
	"go.uber.org/zap"
)

type Trip struct {
	*Tx
	cli    *ent.Client
	logger *zap.SugaredLogger
}

func NewTrip(cli *ent.Client, logger *zap.SugaredLogger) *Trip {
	return &Trip{
		Tx:     &Tx{cli: cli},
		cli:    cli,
		logger: logger.Named("repo.trip"),
	}
}

func (r *Trip) CreateTrip(ctx context.Context, trip *ent.Trip, tx ...*ent.Tx) (*ent.Trip, error) {
	cli := r.cli.Trip
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].Trip
	}

	return cli.Create().
		SetUserID(trip.UserID).
		SetTitle(trip.Title).
		SetDescription(trip.Description).
		SetStartDate(trip.StartDate).
		SetEndDate(trip.EndDate).
		Save(ctx)
}

func (r *Trip) GetTrip(ctx context.Context, id uuid.UUID, tx ...*ent.Tx) (*ent.Trip, error) {
	cli := r.cli.Trip
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].Trip
	}

	return cli.Get(ctx, id)
}

func (r *Trip) UpdateTrip(ctx context.Context, trip *ent.Trip, tx ...*ent.Tx) (*ent.Trip, error) {
	cli := r.cli.Trip
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].Trip
	}

	return cli.UpdateOneID(trip.ID).
		SetTitle(trip.Title).
		SetDescription(trip.Description).
		SetStartDate(trip.StartDate).
		SetEndDate(trip.EndDate).
		SetUpdatedAt(trip.UpdatedAt).
		Save(ctx)
}

func (r *Trip) DeleteTrip(ctx context.Context, id uuid.UUID, tx ...*ent.Tx) error {
	cli := r.cli.Trip
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].Trip
	}

	return cli.DeleteOneID(id).Exec(ctx)
}

func (r *Trip) ListTrips(ctx context.Context, userId uuid.UUID, tx ...*ent.Tx) ([]*ent.Trip, error) {
	cli := r.cli.Trip
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].Trip
	}

	return cli.Query().Where(trip.UserID(userId)).All(ctx)
}

func (r *Trip) CreateDailyTrip(ctx context.Context, dailyTrip *ent.DailyTrip, tx ...*ent.Tx) (*ent.DailyTrip, error) {
	cli := r.cli.DailyTrip
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].DailyTrip
	}

	return cli.Create().
		SetTripID(dailyTrip.TripID).
		SetDay(dailyTrip.Day).
		SetDate(dailyTrip.Date).
		SetNotes(dailyTrip.Notes).
		Save(ctx)
}

func (r *Trip) GetDailyTrip(ctx context.Context, tripId, dailyId uuid.UUID, tx ...*ent.Tx) (*ent.DailyTrip, error) {
	cli := r.cli.DailyTrip
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].DailyTrip
	}

	return cli.Query().
		Where(dailytrip.ID(dailyId), dailytrip.HasTripWith(trip.ID(tripId))).Only(ctx)
}

func (r *Trip) UpdateDailyTrip(ctx context.Context, dailyTrip *ent.DailyTrip, tx ...*ent.Tx) (*ent.DailyTrip, error) {
	cli := r.cli.DailyTrip
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].DailyTrip
	}

	return cli.UpdateOneID(dailyTrip.ID).
		SetDay(dailyTrip.Day).
		SetDate(dailyTrip.Date).
		SetNotes(dailyTrip.Notes).
		SetUpdatedAt(dailyTrip.UpdatedAt).
		Save(ctx)
}

func (r *Trip) DeleteDailyTrip(ctx context.Context, id uuid.UUID, tx ...*ent.Tx) error {
	cli := r.cli.DailyTrip
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].DailyTrip
	}

	return cli.DeleteOneID(id).Exec(ctx)
}

func (r *Trip) ListDailyTrips(ctx context.Context, tripId uuid.UUID, tx ...*ent.Tx) ([]*ent.DailyTrip, error) {
	cli := r.cli.DailyTrip
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].DailyTrip
	}

	return cli.Query().Where(dailytrip.TripID(tripId)).All(ctx)
}
