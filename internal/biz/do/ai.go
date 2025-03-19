package do

import "time"

type (
	PlanAgentInput struct {
		Destination string
		StartDate   time.Time
		EndDate     time.Time
		Duration    int
		Preferences string
		Budget      float64
	}
	PlanAgentOutput struct {
		DailyPlans []*DailyPlan
	}
	DailyPlan struct {
		Day        int
		Date       time.Time
		Title      string
		Activities []*Activity
	}
	Activity struct {
		Time     time.Time
		Name     string
		Duration time.Duration
		Notes    string
	}
)
