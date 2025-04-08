package do

import (
	"time"

	"github.com/google/uuid"
)

type (
	CityPlannerInput struct {
		Destination string
		StartDate   time.Time
		EndDate     time.Time
		Duration    int
		Preferences string
		Budget      float64
	}
	CityPlannerOutput [][]string
)

type (
	TripPlannerInput struct {
		Destination string
		StartDate   time.Time
		EndDate     time.Time
		Duration    int
		Preferences string
		Budget      float64
		POIs        []*PointsOfInterest
	}
	TripPlannerOutput struct {
		Title            string
		Description      string
		StartDate        time.Time
		EndDate          time.Time
		TotalDays        int
		DailyItineraries []*TripPlannerOutputDailyItinerary
	}
	TripPlannerOutputDailyItinerary struct {
		Day   int
		Date  time.Time
		Notes string
		POIs  []*TripPlannerOutputDailyPlanPoi
	}
	TripPlannerOutputDailyPlanPoi struct {
		Id    uuid.UUID
		Notes string
	}
)
