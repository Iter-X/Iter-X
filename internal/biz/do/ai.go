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
	TripPlannerOutput []*DailyPlan
	DailyPlan         struct {
		Day   int
		Date  time.Time
		Title string
		POIs  []*DailyPlanPoi
	}
	DailyPlanPoi struct {
		Id       uuid.UUID
		Time     time.Time
		Name     string
		Duration time.Duration
		Notes    string
	}
)
