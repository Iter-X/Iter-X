package trip

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"text/template"
	"time"

	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/biz/ai/core"
	"github.com/iter-x/iter-x/internal/biz/ai/tool"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/common/cnst"
)

type (
	// tripPlanner is an intelligent travel planning agent
	tripPlanner struct {
		*core.BaseAgent
		toolHub *tool.Hub
	}

	dayPlan struct {
		Day   int          `json:"day"`
		Title string       `json:"title"`
		POIs  []dayPlanPoi `json:"pois"`
	}
	dayPlanPoi struct {
		Time     string `json:"time"`
		ID       string `json:"id"`
		Duration int    `json:"duration"`
		Notes    string `json:"notes"`
	}

	userPromptTpl struct {
		Destination string
		StartDate   time.Time
		EndDate     time.Time
		Duration    int
		Preferences string
		Budget      float64
		POIs        []briefPOI
	}

	briefPOI struct {
		ID   string
		Name string
	}
)

// NewTripPlanner creates a new trip planner agent
func NewTripPlanner(name string, toolHub *tool.Hub, prompt core.Prompt) core.Agent {
	return &tripPlanner{
		BaseAgent: core.NewBaseAgent(name, prompt),
		toolHub:   toolHub,
	}
}

func getPrompt(prompt core.Prompt, input *do.TripPlannerInput, pois []*do.PointsOfInterest) (*do.ToolCompletionInput, error) {
	systemPrompt := prompt.GetSystemPrompt()
	userPrompt := prompt.GetUserPrompt()
	tpl := userPromptTpl{
		Destination: input.Destination,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		Duration:    input.Duration,
		Preferences: input.Preferences,
		Budget:      input.Budget,
	}

	for _, poi := range pois {
		tpl.POIs = append(tpl.POIs, briefPOI{
			ID:   poi.ID.String(),
			Name: poi.NameEn,
		})
	}

	tmpl, err := template.New("user_prompt").Parse(userPrompt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user prompt template: %v", err)
	}

	var userPromptBuf bytes.Buffer
	if err := tmpl.Execute(&userPromptBuf, tpl); err != nil {
		return nil, fmt.Errorf("failed to execute user prompt template: %v", err)
	}

	return &do.ToolCompletionInput{
		Messages: []do.ToolCompletionInputMessage{
			{
				Role:    do.CompletionRoleSystem,
				Content: systemPrompt,
			},
			{
				Role:    do.CompletionRoleUser,
				Content: userPromptBuf.String(),
			},
		},
	}, nil
}

func getItineraries(ctx context.Context, completionTool core.Tool, prompt core.Prompt, input *do.TripPlannerInput, pois []*do.PointsOfInterest) ([]dayPlan, error) {
	completionInput, err := getPrompt(prompt, input, pois)
	if err != nil {
		return nil, err
	}
	resp, err := completionTool.Execute(ctx, completionInput)
	if err != nil {
		return nil, err
	}

	// Extract the raw trip plan from the completion tool
	completionOutput, ok := resp.(*do.ToolCompletionOutput)
	if !ok {
		return nil, fmt.Errorf("invalid completion response type: %T", resp)
	}

	var itineraries []dayPlan
	return itineraries, json.Unmarshal([]byte(completionOutput.Content), &itineraries)
}

// Execute implements the main logic of PlanAgent
func (a *tripPlanner) Execute(ctx context.Context, inputAny any) (any, error) {
	input, ok := inputAny.(*do.TripPlannerInput)
	if !ok {
		return nil, fmt.Errorf("invalid input type: %T", inputAny)
	}

	// get completion tool
	completionTool, err := a.toolHub.GetTool(cnst.ToolCompletion)
	if err != nil {
		return nil, err
	}

	// refine trip plan
	itineraries, err := getItineraries(ctx, completionTool, a.GetPrompt(), input, input.POIs)
	if err != nil {
		return nil, err
	}

	// Convert to TripPlannerOutput
	result := do.TripPlannerOutput{}
	for _, itinerary := range itineraries {
		dp := &do.DailyPlan{
			Day:   itinerary.Day,
			Title: itinerary.Title,
		}

		// Calculate date based on start date and day number
		date := input.StartDate.AddDate(0, 0, itinerary.Day-1)
		dp.Date = date

		for _, poi := range itinerary.POIs {
			// Parse time string to time.Time
			timeStr := fmt.Sprintf("%s %s", date.Format("2006-01-02"), poi.Time)
			activityTime, err := time.Parse("2006-01-02 15:04", timeStr)
			if err != nil {
				return nil, fmt.Errorf("failed to parse time: %v", err)
			}

			poiID, err := uuid.Parse(poi.ID)
			if err != nil {
				return nil, fmt.Errorf("failed to parse DailyPlanPoi ID: %v", err)
			}

			dp.POIs = append(dp.POIs, &do.DailyPlanPoi{
				Id:       poiID,
				Time:     activityTime,
				Name:     poi.Notes,
				Duration: time.Duration(poi.Duration) * time.Second,
				Notes:    poi.Notes,
			})
		}
		result = append(result, dp)
	}

	return &result, nil
}
