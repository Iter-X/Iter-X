package repository

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type DailyItinerary[T *ent.DailyItinerary, R *do.DailyItinerary] interface {
	BaseRepo[T, R]
}

type DailyItineraryRepo = DailyItinerary[*ent.DailyItinerary, *do.DailyItinerary]
