package plan

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"text/template"
	"time"

	"github.com/google/uuid"
	"github.com/ifuryst/lol"

	"github.com/iter-x/iter-x/internal/biz/ai/core"
	"github.com/iter-x/iter-x/internal/biz/ai/tool"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/common/cnst"
)

type (
	// agent is an intelligent travel planning agent
	agent struct {
		*core.BaseAgent
		toolHub *tool.Hub
		poiRepo repository.PointsOfInterestRepo
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

	r1UserPromptTpl struct {
		Destination string
		StartDate   time.Time
		EndDate     time.Time
		Duration    int
		Preferences string
		Budget      float64
	}

	r2UserPromptTpl struct {
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

// NewAgent creates a new PlanAgent
func NewAgent(name string, toolHub *tool.Hub, prompt core.Prompt, poiRepo repository.PointsOfInterestRepo) core.Agent {
	return &agent{
		BaseAgent: core.NewBaseAgent(name, prompt),
		toolHub:   toolHub,
		poiRepo:   poiRepo,
	}
}

func getRound1Prompt(prompt core.Prompt, input *do.PlanAgentInput) (*do.ToolCompletionInput, error) {
	systemPrompt := prompt.GetSystemPrompt()
	userPrompt := prompt.GetUserPrompt()
	tmplData := r1UserPromptTpl{
		Destination: input.Destination,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		Duration:    input.Duration,
		Preferences: input.Preferences,
		Budget:      input.Budget,
	}
	tmpl, err := template.New("round1_user_prompt").Parse(userPrompt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse round1 user prompt template: %v", err)
	}

	var userPromptBuf bytes.Buffer
	if err := tmpl.Execute(&userPromptBuf, tmplData); err != nil {
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

func getPotentialCities(ctx context.Context, completionTool core.Tool, prompt core.Prompt, input *do.PlanAgentInput) ([][]string, error) {
	completionInput, err := getRound1Prompt(prompt, input)
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

	var cities [][]string
	return cities, json.Unmarshal([]byte(completionOutput.Content), &cities)
}

func getRound2Prompt(prompt core.Prompt, input *do.PlanAgentInput, pois []*do.PointsOfInterest) (*do.ToolCompletionInput, error) {
	systemPrompt := prompt.GetSystemPromptByRound(2)
	userPrompt := prompt.GetUserPromptByRound(2)
	tpl := r2UserPromptTpl{
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

	tmpl, err := template.New("round2_user_prompt").Parse(userPrompt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse round2 user prompt template: %v", err)
	}

	var userPromptBuf bytes.Buffer
	if err := tmpl.Execute(&userPromptBuf, tpl); err != nil {
		return nil, fmt.Errorf("failed to execute round2 user prompt template: %v", err)
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

func getItineraries(ctx context.Context, completionTool core.Tool, prompt core.Prompt, input *do.PlanAgentInput, pois []*do.PointsOfInterest) ([]dayPlan, error) {
	completionInput, err := getRound2Prompt(prompt, input, pois)
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
func (a *agent) Execute(ctx context.Context, inputAny any) (any, error) {
	input, ok := inputAny.(*do.PlanAgentInput)
	if !ok {
		return nil, fmt.Errorf("invalid input type: %T", inputAny)
	}

	// get completion tool
	completionTool, err := a.toolHub.GetTool(cnst.ToolCompletion)
	if err != nil {
		return nil, err
	}

	// Round 1: get potential cities
	potentialCities, err := getPotentialCities(ctx, completionTool, a.GetPrompt(), input)
	if err != nil {
		return nil, err
	}

	// Retrieve POIs for potential cities
	pois, err := a.poiRepo.GetByCityNames(ctx, lol.UniqSlice(potentialCities...))
	if err != nil {
		return nil, fmt.Errorf("failed to get POIs for potential cities: %v", err)
	}

	// Round 2: refine trip plan
	itineraries, err := getItineraries(ctx, completionTool, a.GetPrompt(), input, pois)
	if err != nil {
		return nil, err
	}

	// Convert to PlanAgentOutput
	result := do.PlanAgentOutput{}
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
