package do

import (
	"time"

	"github.com/google/uuid"
)

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

	User *User `json:"user,omitempty"`
}

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
	// DailyTripItem holds the value of the daily_trip_item edge.
	DailyTripItem []*DailyTripItem `json:"daily_trip_item,omitempty"`
}

type DailyTripItem struct {
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
	// Notes holds the value of the "notes" field.
	Notes string `json:"notes,omitempty"`

	// Trip holds the value of the trip edge.
	Trip *Trip `json:"trip,omitempty"`
	// DailyTrip holds the value of the daily_trip edge.
	DailyTrip *DailyTrip `json:"daily_trip,omitempty"`
}
