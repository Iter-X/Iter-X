package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/impl/ent"
)

// Trip .
type Trip[T *ent.Trip, R *do.Trip] interface {
	Base[T, R]

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
}

type TripRepo = Trip[*ent.Trip, *do.Trip]

// DailyTrip .
type DailyTrip[T *ent.DailyTrip, R *do.DailyTrip] interface {
	Base[T, R]
}

type DailyTripRepo = DailyTrip[*ent.DailyTrip, *do.DailyTrip]

// DailyTripItem .
type DailyTripItem[T *ent.DailyTripItem, R *do.DailyTripItem] interface {
	Base[T, R]
}

type DailyTripItemRepo = DailyTripItem[*ent.DailyTripItem, *do.DailyTripItem]
