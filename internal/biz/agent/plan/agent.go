package plan

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/agent/core"
)

// Agent is an intelligent travel planning agent
type Agent struct {
	*core.BaseAgent
}

// NewAgent creates a new PlanAgent
func NewAgent(name, desc string, tools []core.Tool, prompt core.Prompt) *Agent {
	return &Agent{
		BaseAgent: core.NewBaseAgent(
			name,
			desc,
			tools,
			prompt,
		),
	}
}

// Execute implements the main logic of PlanAgent
func (a *Agent) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	// Implement the specific travel planning logic here
	// 1. Parse input parameters
	// 2. Use the toolset for planning
	// 3. Generate travel plan
	// 4. Return the result

	return nil, nil
}

// Input defines the input parameter structure for PlanAgent
type Input struct {
	Destination string   `json:"destination"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Preferences []string `json:"preferences"`
	Budget      float64  `json:"budget"`
}

// Output defines the output result structure for PlanAgent
type Output struct {
	Itinerary []DayPlan `json:"itinerary"`
	Summary   string    `json:"summary"`
}

// DayPlan represents a day's travel plan
type DayPlan struct {
	Date     string     `json:"date"`
	Schedule []Activity `json:"schedule"`
}

// Activity represents a specific activity
type Activity struct {
	Time        string  `json:"time"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Duration    int     `json:"duration"` // minutes
	Cost        float64 `json:"cost"`
}
