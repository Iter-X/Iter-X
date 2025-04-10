package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type Trip[T *ent.Trip, R *do.Trip] interface {
	BaseRepo[T, R]

	CreateTrip(ctx context.Context, trip *do.Trip) (*do.Trip, error)

	GetTrip(ctx context.Context, id uuid.UUID) (*do.Trip, error)

	UpdateTrip(ctx context.Context, trip *do.Trip) (*do.Trip, error)

	DeleteTrip(ctx context.Context, id uuid.UUID) error

	ListTrips(ctx context.Context, userId uuid.UUID) ([]*do.Trip, error)

	CreateDailyTrip(ctx context.Context, dailyTrip *do.DailyTrip) (*do.DailyTrip, error)

	GetDailyTrip(ctx context.Context, tripId, dailyId uuid.UUID) (*do.DailyTrip, error)

	UpdateDailyTrip(ctx context.Context, dailyTrip *do.DailyTrip) (*do.DailyTrip, error)

	DeleteDailyTrip(ctx context.Context, id uuid.UUID) error

	ListDailyTrips(ctx context.Context, tripId uuid.UUID) ([]*do.DailyTrip, error)

	CreateDailyItinerary(ctx context.Context, dailyItinerary *do.DailyItinerary) (*do.DailyItinerary, error)

	ListTripCollaborators(ctx context.Context, tripId uuid.UUID) ([]*do.TripCollaborator, error)

	AddTripCollaborators(ctx context.Context, tripId uuid.UUID, userIds []uuid.UUID) ([]*do.TripCollaborator, error)

	RemoveTripCollaborator(ctx context.Context, tripId uuid.UUID, userId uuid.UUID) error

	UpdateCollaboratorStatus(ctx context.Context, tripId uuid.UUID, userId uuid.UUID, status string) (*do.TripCollaborator, error)

	CreateTripPOIPool(ctx context.Context, tripPOIPool *do.TripPOIPool) (*do.TripPOIPool, error)

	DeleteTripPOIPool(ctx context.Context, id uuid.UUID) error

	ListTripPOIPool(ctx context.Context, tripId uuid.UUID) ([]*do.TripPOIPool, error)
}

type TripRepo = Trip[*ent.Trip, *do.Trip]

type DailyTrip[T *ent.DailyTrip, R *do.DailyTrip] interface {
	BaseRepo[T, R]
}

type DailyTripRepo = DailyTrip[*ent.DailyTrip, *do.DailyTrip]

type TripPOIPool[T *ent.TripPOIPool, R *do.TripPOIPool] interface {
	BaseRepo[T, R]

	CreateTripPOIPool(ctx context.Context, tripPOIPool *do.TripPOIPool) (*do.TripPOIPool, error)

	DeleteTripPOIPool(ctx context.Context, id uuid.UUID) error

	ListTripPOIPool(ctx context.Context, tripId uuid.UUID) ([]*do.TripPOIPool, error)
}

type TripPOIPoolRepo = TripPOIPool[*ent.TripPOIPool, *do.TripPOIPool]
