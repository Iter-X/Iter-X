package do

import (
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	tripV1 "github.com/iter-x/iter-x/internal/api/trip/v1"
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

	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// DailyTrip holds the value of the daily_trip edge.
	DailyTrip []*DailyTrip `json:"daily_trip,omitempty"`
	// DailyItinerary holds the value of the daily_itinerary edge.
	DailyItinerary []*DailyItinerary `json:"daily_itinerary,omitempty"`
}

func (t *Trip) ToTripProto() *tripV1.Trip {
	return &tripV1.Trip{
		Id:        t.ID.String(),
		Status:    t.Status,
		Title:     t.Title,
		StartTs:   timestamppb.New(t.StartDate),
		EndTs:     timestamppb.New(t.EndDate),
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
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
}

func (dt *DailyTrip) ToDailyTripProto() *tripV1.DailyTrip {
	return &tripV1.DailyTrip{
		Id:        dt.ID.String(),
		TripId:    dt.TripID.String(),
		Day:       dt.Day,
		Date:      timestamppb.New(dt.Date),
		Notes:     dt.Notes,
		CreatedAt: timestamppb.New(dt.CreatedAt),
		UpdatedAt: timestamppb.New(dt.UpdatedAt),
	}
}
