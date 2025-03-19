package bo

import (
	"github.com/iter-x/iter-x/internal/common/cnst"
	"time"

	"github.com/google/uuid"
)

type (
	Trip struct {
		ID          uuid.UUID
		UserID      uuid.UUID
		Title       string
		Description string
		StartDate   time.Time
		EndDate     time.Time
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Status      bool
	}

	CreateTripRequest struct {
		CreationMethod cnst.TripCreationMethod
		Destination    string
		StartDate      time.Time
		EndDate        time.Time
		Duration       int
	}
)

type ListDailyTripsRequest struct {
	TripID string `json:"trip_id"`
}

type DeleteDailyTripRequest struct {
	TripID  string `json:"trip_id"`
	DailyID string `json:"daily_id"`
}

type UpdateDailyTripRequest struct {
	TripID  string    `json:"trip_id"`
	DailyID string    `json:"daily_id"`
	Day     int32     `json:"day"`
	Date    time.Time `json:"date"`
	Notes   string    `json:"notes"`
}

type GetDailyTripRequest struct {
	TripID  string `json:"trip_id"`
	DailyID string `json:"daily_id"`
}

type CreateDailyTripRequest struct {
	TripID string    `json:"trip_id"`
	Day    int32     `json:"day"`
	Date   time.Time `json:"date"`
	Notes  string    `json:"notes"`
}

type UpdateTripRequest struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Status      bool      `json:"status"`
}
