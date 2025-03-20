package bo

import (
	"time"
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

type CreateTripRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}
