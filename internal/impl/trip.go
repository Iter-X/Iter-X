package impl

import (
	"context"

	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/impl/ent"
	"github.com/iter-x/iter-x/internal/impl/ent/dailytrip"
	"github.com/iter-x/iter-x/internal/impl/ent/trip"
	"go.uber.org/zap"
)

func NewTripRepository(cli *ent.Client, logger *zap.SugaredLogger) repository.Trip[*ent.Trip, *do.Trip] {
	return &tripRepositoryImpl{
		Tx:                      &Tx{cli: cli},
		logger:                  logger.Named("repo.trip"),
		authRepositoryImpl:      new(authRepositoryImpl),
		tripDailyRepositoryImpl: new(tripDailyRepositoryImpl),
	}
}

type tripRepositoryImpl struct {
	*Tx
	logger                  *zap.SugaredLogger
	authRepositoryImpl      repository.Base[*ent.User, *do.User]
	tripDailyRepositoryImpl repository.Base[*ent.DailyTrip, *do.DailyTrip]
}

func (r *tripRepositoryImpl) ToEntity(po *ent.Trip) *do.Trip {
	if po == nil {
		return nil
	}
	return &do.Trip{
		ID:          po.ID,
		CreatedAt:   po.CreatedAt,
		UpdatedAt:   po.UpdatedAt,
		UserID:      po.UserID,
		Status:      po.Status,
		Title:       po.Title,
		Description: po.Description,
		StartDate:   po.StartDate,
		EndDate:     po.EndDate,
		User:        r.authRepositoryImpl.ToEntity(po.Edges.User),
	}
}

func (r *tripRepositoryImpl) ToEntities(pos []*ent.Trip) []*do.Trip {
	list := make([]*do.Trip, 0, len(pos))
	for _, v := range pos {
		list = append(list, r.ToEntity(v))
	}
	return list
}

func (r *tripRepositoryImpl) CreateTrip(ctx context.Context, trip *do.Trip) (*do.Trip, error) {
	cli := r.GetTx(ctx).Trip

	row, err := cli.Create().
		SetUserID(trip.UserID).
		SetTitle(trip.Title).
		SetDescription(trip.Description).
		SetStartDate(trip.StartDate).
		SetEndDate(trip.EndDate).
		Save(ctx)
	return r.ToEntity(row), err
}

func (r *tripRepositoryImpl) GetTrip(ctx context.Context, id uuid.UUID) (*do.Trip, error) {
	cli := r.GetTx(ctx).Trip
	row, err := cli.Get(ctx, id)
	return r.ToEntity(row), err
}

func (r *tripRepositoryImpl) UpdateTrip(ctx context.Context, trip *do.Trip) (*do.Trip, error) {
	cli := r.GetTx(ctx).Trip

	row, err := cli.UpdateOneID(trip.ID).
		SetTitle(trip.Title).
		SetDescription(trip.Description).
		SetStartDate(trip.StartDate).
		SetEndDate(trip.EndDate).
		SetUpdatedAt(trip.UpdatedAt).
		Save(ctx)
	return r.ToEntity(row), err
}

func (r *tripRepositoryImpl) DeleteTrip(ctx context.Context, id uuid.UUID) error {
	cli := r.GetTx(ctx).Trip

	return cli.DeleteOneID(id).Exec(ctx)
}

func (r *tripRepositoryImpl) ListTrips(ctx context.Context, userId uuid.UUID) ([]*ent.Trip, error) {
	cli := r.GetTx(ctx).Trip

	return cli.Query().Where(trip.UserID(userId)).All(ctx)
}

func (r *tripRepositoryImpl) CreateDailyTrip(ctx context.Context, dailyTrip *do.DailyTrip) (*do.DailyTrip, error) {
	cli := r.GetTx(ctx).DailyTrip

	row, err := cli.Create().
		SetTripID(dailyTrip.TripID).
		SetDay(dailyTrip.Day).
		SetDate(dailyTrip.Date).
		SetNotes(dailyTrip.Notes).
		Save(ctx)
	return r.tripDailyRepositoryImpl.ToEntity(row), err
}

func (r *tripRepositoryImpl) GetDailyTrip(ctx context.Context, tripId, dailyId uuid.UUID) (*do.DailyTrip, error) {
	cli := r.GetTx(ctx).DailyTrip

	row, err := cli.Query().
		Where(dailytrip.ID(dailyId), dailytrip.HasTripWith(trip.ID(tripId))).Only(ctx)
	return r.tripDailyRepositoryImpl.ToEntity(row), err
}

func (r *tripRepositoryImpl) UpdateDailyTrip(ctx context.Context, dailyTrip *do.DailyTrip) (*do.DailyTrip, error) {
	cli := r.GetTx(ctx).DailyTrip

	row, err := cli.UpdateOneID(dailyTrip.ID).
		SetDay(dailyTrip.Day).
		SetDate(dailyTrip.Date).
		SetNotes(dailyTrip.Notes).
		SetUpdatedAt(dailyTrip.UpdatedAt).
		Save(ctx)
	return r.tripDailyRepositoryImpl.ToEntity(row), err
}

func (r *tripRepositoryImpl) DeleteDailyTrip(ctx context.Context, id uuid.UUID) error {
	cli := r.GetTx(ctx).DailyTrip

	return cli.DeleteOneID(id).Exec(ctx)
}

func (r *tripRepositoryImpl) ListDailyTrips(ctx context.Context, tripId uuid.UUID) ([]*do.DailyTrip, error) {
	cli := r.GetTx(ctx).DailyTrip

	rows, err := cli.Query().Where(dailytrip.TripID(tripId)).All(ctx)
	return r.tripDailyRepositoryImpl.ToEntities(rows), err
}
