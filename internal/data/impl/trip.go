package impl

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/cnst"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/dailyitinerary"
	"github.com/iter-x/iter-x/internal/data/ent/dailytrip"
	"github.com/iter-x/iter-x/internal/data/ent/trip"
	"github.com/iter-x/iter-x/internal/data/ent/tripcollaborator"
	"github.com/iter-x/iter-x/internal/data/ent/trippoipool"
	"github.com/iter-x/iter-x/internal/data/impl/build"
)

func NewTrip(d *data.Data, logger *zap.SugaredLogger) repository.TripRepo {
	return &tripRepositoryImpl{
		Tx:     d.Tx,
		logger: logger.Named("repo.trip"),
	}
}

type tripRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger
}

func (r *tripRepositoryImpl) ToEntity(po *ent.Trip) *do.Trip {
	if po == nil {
		return nil
	}

	return build.TripRepositoryImplToEntity(po)
}

func (r *tripRepositoryImpl) ToEntities(pos []*ent.Trip) []*do.Trip {
	if pos == nil {
		return nil
	}

	return build.TripRepositoryImplToEntities(pos)
}

func (r *tripRepositoryImpl) CreateTrip(ctx context.Context, trip *do.Trip) (*do.Trip, error) {
	cli := r.GetTx(ctx).Trip

	row, err := cli.Create().
		SetUserID(trip.UserID).
		SetTitle(trip.Title).
		SetDescription(trip.Description).
		SetStartDate(trip.StartDate).
		SetEndDate(trip.EndDate).
		SetDays(trip.Days).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// Add creator as the first collaborator with 'accepted' status
	_, err = r.GetTx(ctx).TripCollaborator.Create().
		SetTripID(row.ID).
		SetUserID(trip.UserID).
		SetStatus(cnst.CollaboratorStatusAccepted).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.ToEntity(row), nil
}

func (r *tripRepositoryImpl) GetTrip(ctx context.Context, id uuid.UUID) (*do.Trip, error) {
	cli := r.GetTx(ctx).Trip

	row, err := cli.Query().
		Where(trip.ID(id)).
		WithDailyTrip(func(q *ent.DailyTripQuery) {
			q.Order(ent.Asc(dailytrip.FieldDay))
			q.WithDailyItinerary(func(q *ent.DailyItineraryQuery) {
				q.Order(ent.Asc("order")).
					WithPoi()
			})
		}).
		Only(ctx)
	return r.ToEntity(row), err
}

func (r *tripRepositoryImpl) UpdateTrip(ctx context.Context, trip *do.Trip) (*do.Trip, error) {
	cli := r.GetTx(ctx).Trip

	row, err := cli.UpdateOneID(trip.ID).
		SetTitle(trip.Title).
		SetDescription(trip.Description).
		SetStartDate(trip.StartDate).
		SetEndDate(trip.EndDate).
		SetDays(trip.Days).
		SetUpdatedAt(trip.UpdatedAt).
		Save(ctx)
	return r.ToEntity(row), err
}

func (r *tripRepositoryImpl) DeleteTrip(ctx context.Context, id uuid.UUID) error {
	cli := r.GetTx(ctx).Trip

	return cli.DeleteOneID(id).Exec(ctx)
}

func (r *tripRepositoryImpl) ListTrips(ctx context.Context, userId uuid.UUID) ([]*do.Trip, error) {
	cli := r.GetTx(ctx).Trip

	rows, err := cli.Query().Where(trip.UserID(userId)).All(ctx)
	return r.ToEntities(rows), err
}

func (r *tripRepositoryImpl) CreateDailyTrip(ctx context.Context, dailyTrip *do.DailyTrip) (*do.DailyTrip, error) {
	cli := r.GetTx(ctx).DailyTrip

	row, err := cli.Create().
		SetTripID(dailyTrip.TripID).
		SetDay(dailyTrip.Day).
		SetDate(dailyTrip.Date).
		SetNotes(dailyTrip.Notes).
		Save(ctx)
	return build.DailyTripRepositoryImplToEntity(row), err
}

func (r *tripRepositoryImpl) GetDailyTrip(ctx context.Context, tripId, dailyId uuid.UUID) (*do.DailyTrip, error) {
	cli := r.GetTx(ctx).DailyTrip

	row, err := cli.Query().
		Where(dailytrip.ID(dailyId), dailytrip.HasTripWith(trip.ID(tripId))).Only(ctx)
	return build.DailyTripRepositoryImplToEntity(row), err
}

func (r *tripRepositoryImpl) UpdateDailyTrip(ctx context.Context, dailyTrip *do.DailyTrip) (*do.DailyTrip, error) {
	cli := r.GetTx(ctx).DailyTrip

	row, err := cli.UpdateOneID(dailyTrip.ID).
		SetDay(dailyTrip.Day).
		SetDate(dailyTrip.Date).
		SetNotes(dailyTrip.Notes).
		SetUpdatedAt(dailyTrip.UpdatedAt).
		Save(ctx)
	return build.DailyTripRepositoryImplToEntity(row), err
}

func (r *tripRepositoryImpl) DeleteDailyTrip(ctx context.Context, id uuid.UUID) error {
	cli := r.GetTx(ctx).DailyTrip

	return cli.DeleteOneID(id).Exec(ctx)
}

func (r *tripRepositoryImpl) ListDailyTrips(ctx context.Context, tripId uuid.UUID) ([]*do.DailyTrip, error) {
	cli := r.GetTx(ctx).DailyTrip

	rows, err := cli.Query().Where(dailytrip.TripID(tripId)).All(ctx)
	return build.DailyTripRepositoryImplToEntities(rows), err
}

func (r *tripRepositoryImpl) CreateDailyItinerary(ctx context.Context, dailyItinerary *do.DailyItinerary) (*do.DailyItinerary, error) {
	cli := r.GetTx(ctx).DailyItinerary

	row, err := cli.Create().
		SetCreatedAt(dailyItinerary.CreatedAt).
		SetUpdatedAt(dailyItinerary.UpdatedAt).
		SetTripID(dailyItinerary.TripID).
		SetDailyTripID(dailyItinerary.DailyTripID).
		SetPoiID(dailyItinerary.PoiID).
		SetNotes(dailyItinerary.Notes).
		SetOrder(dailyItinerary.Order).
		Save(ctx)
	return build.DailyItineraryRepositoryImplToEntity(row), err
}

func (r *tripRepositoryImpl) GetDailyItinerary(ctx context.Context, tripId, dailyTripId, dailyItineraryId uuid.UUID) (*do.DailyItinerary, error) {
	cli := r.GetTx(ctx).DailyItinerary

	row, err := cli.Query().
		Where(
			dailyitinerary.ID(dailyItineraryId),
			dailyitinerary.TripID(tripId),
			dailyitinerary.DailyTripID(dailyTripId),
		).
		WithPoi().
		Only(ctx)
	return build.DailyItineraryRepositoryImplToEntity(row), err
}

func (r *tripRepositoryImpl) ListDailyItinerariesByDay(ctx context.Context, tripId uuid.UUID, day int32) ([]*do.DailyItinerary, error) {
	cli := r.GetTx(ctx).DailyItinerary

	rows, err := cli.Query().
		Where(
			dailyitinerary.HasDailyTripWith(
				dailytrip.TripID(tripId),
				dailytrip.Day(day),
			),
		).
		Order(ent.Asc(dailyitinerary.FieldOrder)).
		WithPoi().
		All(ctx)
	return build.DailyItineraryRepositoryImplToEntities(rows), err
}

func (r *tripRepositoryImpl) UpdateDailyItinerary(ctx context.Context, dailyItinerary *do.DailyItinerary) (*do.DailyItinerary, error) {
	cli := r.GetTx(ctx).DailyItinerary

	row, err := cli.UpdateOneID(dailyItinerary.ID).
		SetOrder(dailyItinerary.Order).
		SetNotes(dailyItinerary.Notes).
		SetUpdatedAt(dailyItinerary.UpdatedAt).
		Save(ctx)
	return build.DailyItineraryRepositoryImplToEntity(row), err
}

func (r *tripRepositoryImpl) DeleteDailyItinerary(ctx context.Context, id uuid.UUID) error {
	cli := r.GetTx(ctx).DailyItinerary

	return cli.DeleteOneID(id).Exec(ctx)
}

func (r *tripRepositoryImpl) ListTripCollaborators(ctx context.Context, tripId uuid.UUID) ([]*do.TripCollaborator, error) {
	cli := r.GetTx(ctx).Trip

	// Get all collaborators through trip_collaborators with user information
	collaborators, err := cli.Query().
		Where(trip.ID(tripId)).
		QueryTripCollaborators().
		WithUser().
		All(ctx)
	if err != nil {
		return nil, err
	}

	return build.TripCollaboratorRepositoryImplToEntities(collaborators), nil
}

func (r *tripRepositoryImpl) CreateTripPOIPool(ctx context.Context, tripPOIPool *do.TripPOIPool) (*do.TripPOIPool, error) {
	cli := r.GetTx(ctx).TripPOIPool

	row, err := cli.Create().
		SetCreatedAt(tripPOIPool.CreatedAt).
		SetUpdatedAt(tripPOIPool.UpdatedAt).
		SetTripID(tripPOIPool.TripID).
		SetPoiID(tripPOIPool.PoiID).
		SetCreatedBy(tripPOIPool.CreatedBy).
		SetUpdatedBy(tripPOIPool.UpdatedBy).
		Save(ctx)
	return build.TripPOIPoolRepositoryImplToEntity(row), err
}

func (r *tripRepositoryImpl) DeleteTripPOIPool(ctx context.Context, id uuid.UUID) error {
	cli := r.GetTx(ctx).TripPOIPool

	return cli.DeleteOneID(id).Exec(ctx)
}

func (r *tripRepositoryImpl) ListTripPOIPool(ctx context.Context, tripId uuid.UUID) ([]*do.TripPOIPool, error) {
	cli := r.GetTx(ctx).TripPOIPool

	rows, err := cli.Query().
		Where(trippoipool.TripID(tripId)).
		WithPoi().
		All(ctx)
	return build.TripPOIPoolRepositoryImplToEntities(rows), err
}

func (r *tripRepositoryImpl) AddTripCollaborators(ctx context.Context, tripId uuid.UUID, userIds []uuid.UUID) ([]*do.TripCollaborator, error) {
	cli := r.GetTx(ctx).TripCollaborator

	var collaborators []*ent.TripCollaborator
	for _, userId := range userIds {
		// Check if the user is already a collaborator
		exists, err := cli.Query().
			Where(
				tripcollaborator.TripID(tripId),
				tripcollaborator.UserID(userId),
			).
			Exist(ctx)
		if err != nil {
			return nil, err
		}
		if exists {
			continue
		}

		// Add new collaborator if not exists
		collaborator, err := cli.Create().
			SetTripID(tripId).
			SetUserID(userId).
			SetStatus(cnst.CollaboratorStatusInvited).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		collaborators = append(collaborators, collaborator)
	}

	return build.TripCollaboratorRepositoryImplToEntities(collaborators), nil
}

func (r *tripRepositoryImpl) RemoveTripCollaborator(ctx context.Context, tripId uuid.UUID, userId uuid.UUID) error {
	cli := r.GetTx(ctx).TripCollaborator

	_, err := cli.Delete().
		Where(
			tripcollaborator.TripID(tripId),
			tripcollaborator.UserID(userId),
		).
		Exec(ctx)
	return err
}

func (r *tripRepositoryImpl) UpdateCollaboratorStatus(ctx context.Context, tripId uuid.UUID, userId uuid.UUID, status string) (*do.TripCollaborator, error) {
	cli := r.GetTx(ctx).TripCollaborator

	collaborator, err := cli.Query().
		Where(
			tripcollaborator.TripID(tripId),
			tripcollaborator.UserID(userId),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	updated, err := collaborator.Update().
		SetStatus(tripcollaborator.Status(status)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return build.TripCollaboratorRepositoryImplToEntity(updated), nil
}
