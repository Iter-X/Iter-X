package do

import (
	"time"

	"github.com/google/uuid"
)

// DailyItinerary is the model entity for the DailyItinerary schema.
type DailyItinerary struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// TripID holds the value of the "trip_id" field.
	TripID uuid.UUID `json:"trip_id,omitempty"`
	// DailyTripID holds the value of the "daily_trip_id" field.
	DailyTripID uuid.UUID `json:"daily_trip_id,omitempty"`
	// PoiID holds the value of the "poi_id" field.
	PoiID uuid.UUID `json:"poi_id,omitempty"`
	// Notes holds the value of the "notes" field.
	Notes string `json:"notes,omitempty"`

	// Trip holds the value of the trip edge.
	Trip *Trip `json:"trip,omitempty"`
	// DailyTrip holds the value of the daily_trip edge.
	DailyTrip *DailyTrip `json:"daily_trip,omitempty"`
	// Poi holds the value of the poi edge.
	Poi *PointsOfInterest `json:"poi,omitempty"`
}
