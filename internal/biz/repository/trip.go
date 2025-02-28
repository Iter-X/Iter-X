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

	CreateDailyTrip(ctx context.Context, dailyTrip *do.DailyTrip) (*do.DailyTrip, error)
}

// DailyTrip .
type DailyTrip[T *ent.DailyTrip, R *do.DailyTrip] interface {
	Base[T, R]
}
