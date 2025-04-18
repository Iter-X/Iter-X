package do

import (
	"time"

	"github.com/google/uuid"
)

// Trip is the model entity for the Trip schema.
type Trip struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// StartDate holds the value of the "start_date" field.
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate holds the value of the "end_date" field.
	EndDate time.Time `json:"end_date,omitempty"`
	// Days holds the value of the "days" field.
	Days int8 `json:"days,omitempty"`

	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// DailyTrip holds the value of the daily_trip edge.
	DailyTrip []*DailyTrip `json:"daily_trip,omitempty"`
	// DailyItinerary holds the value of the daily_itinerary edge.
	DailyItinerary []*DailyItinerary `json:"daily_itinerary,omitempty"`
	// PoiPool holds the value of the poi_pool edge.
	PoiPool []*TripPOIPool `json:"poi_pool,omitempty"`
}

// DailyTrip is the model entity for the DailyTrip schema.
type DailyTrip struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// TripID holds the value of the "trip_id" field.
	TripID uuid.UUID `json:"trip_id,omitempty"`
	// Day holds the value of the "day" field.
	Day int32 `json:"day,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// Notes holds the value of the "notes" field.
	Notes string `json:"notes,omitempty"`

	// Trip holds the value of the trip edge.
	Trip *Trip `json:"trip,omitempty"`
	// DailyItinerary holds the value of the daily_itinerary edge.
	DailyItinerary []*DailyItinerary `json:"daily_itinerary,omitempty"`
}

// TripPOIPool is the model entity for the TripPOIPool schema.
type TripPOIPool struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy uuid.UUID `json:"updated_by,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy uuid.UUID `json:"created_by,omitempty"`
	// TripID holds the value of the "trip_id" field.
	TripID uuid.UUID `json:"trip_id,omitempty"`
	// PoiID holds the value of the "poi_id" field.
	PoiID uuid.UUID `json:"poi_id,omitempty"`

	// Trip holds the value of the trip edge.
	Trip *Trip `json:"trip,omitempty"`
	// Poi holds the value of the poi edge.
	Poi *PointsOfInterest `json:"poi,omitempty"`
}
