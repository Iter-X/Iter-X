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
	"go.uber.org/zap"
)

type (
	// tripPlanner is an intelligent travel planning agent
	tripPlanner struct {
		*core.BaseAgent
		toolHub *tool.Hub
		logger  *zap.SugaredLogger
	}

	createTripArgs struct {
		Title          string               `json:"title"`
		Description    string               `json:"description"`
		StartDate      string               `json:"start_date"`
		EndDate        string               `json:"end_date"`
		TotalDays      int                  `json:"total_days"`
		DailyItinerary []dailyItineraryItem `json:"daily_itinerary"`
	}

	dailyItineraryItem struct {
		Day   int    `json:"day"`
		Date  string `json:"date"`
		Notes string `json:"notes"`
		POIs  []struct {
			ID    string `json:"id"`
			Notes string `json:"notes"`
		} `json:"pois"`
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
		City string
	}
)

// NewTripPlanner creates a new trip planner agent
func NewTripPlanner(name string, toolHub *tool.Hub, prompt core.Prompt, toolNames []string, logger *zap.SugaredLogger) core.Agent {
	agent := &tripPlanner{
		BaseAgent: core.NewBaseAgent(name, prompt, toolNames),
		toolHub:   toolHub,
		logger:    logger.Named("agent.trip_planner"),
	}
	return agent
}

func getPrompt(prompt core.Prompt, input *do.TripPlannerInput, pois []*do.PointsOfInterest,
	toolHub *tool.Hub, toolNames []string) (*do.ToolCompletionInput, error) {
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
			Name: poi.NameCn,
			City: poi.City.GetNameCn(),
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

	// Get all tools from the agent's configuration
	tools := make([]do.FunctionCallTool, 0)
	for _, name := range toolNames {
		t, err := toolHub.GetTool(name)
		if err != nil {
			continue
		}

		if def, err := t.GetDefinition(); err == nil {
			tools = append(tools, *def)
		}
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
		Tools: tools,
	}, nil
}

func getItineraries(ctx context.Context, completionTool core.Tool, prompt core.Prompt,
	input *do.TripPlannerInput, pois []*do.PointsOfInterest, toolHub *tool.Hub,
	toolNames []string, logger *zap.SugaredLogger) (*do.TripPlannerOutput, error) {
	completionInput, err := getPrompt(prompt, input, pois, toolHub, toolNames)
	if err != nil {
		logger.Errorw("failed to get prompt", "err", err)
		return nil, err
	}
	resp, err := completionTool.Execute(ctx, completionInput)
	if err != nil {
		logger.Errorw("failed to execute completion", "err", err)
		return nil, err
	}

	// Extract the raw trip plan from the completion tool
	completionOutput, ok := resp.(*do.ToolCompletionOutput[[]do.ToolCallOutput])
	if !ok {
		logger.Errorw("invalid completion response type", "type", fmt.Sprintf("%T", resp))
		return nil, fmt.Errorf("invalid completion response type: %T", resp)
	}

	// Find create_trip function call
	var createTripCall *do.ToolCallOutput
	for _, call := range completionOutput.Content {
		if call.Function.Name == "create_trip" {
			createTripCall = &call
			break
		}
	}
	if createTripCall == nil {
		logger.Errorw("create_trip function call not found in response")
		return nil, fmt.Errorf("create_trip function call not found in response")
	}

	// Parse create_trip arguments
	var args createTripArgs
	if err := json.Unmarshal([]byte(createTripCall.Function.Arguments), &args); err != nil {
		logger.Errorw("failed to unmarshal create_trip arguments", "err", err)
		return nil, fmt.Errorf("failed to unmarshal create_trip arguments: %v", err)
	}

	// Convert to TripPlannerOutput
	startDate, err := time.Parse("2006-01-02", args.StartDate)
	if err != nil {
		logger.Errorw("failed to parse start date", "err", err)
		return nil, fmt.Errorf("failed to parse start date: %v", err)
	}
	endDate, err := time.Parse("2006-01-02", args.EndDate)
	if err != nil {
		logger.Errorw("failed to parse end date", "err", err)
		return nil, fmt.Errorf("failed to parse end date: %v", err)
	}

	output := &do.TripPlannerOutput{
		Title:            args.Title,
		Description:      args.Description,
		StartDate:        startDate,
		EndDate:          endDate,
		TotalDays:        args.TotalDays,
		DailyItineraries: make([]*do.TripPlannerOutputDailyItinerary, 0, len(args.DailyItinerary)),
	}

	// Convert daily schedules
	for _, schedule := range args.DailyItinerary {
		date, err := time.Parse("2006-01-02", schedule.Date)
		if err != nil {
			logger.Errorw("failed to parse schedule date", "err", err)
			return nil, fmt.Errorf("failed to parse schedule date: %v", err)
		}

		dailyItinerary := &do.TripPlannerOutputDailyItinerary{
			Day:   schedule.Day,
			Date:  date,
			Notes: schedule.Notes,
			POIs:  make([]*do.TripPlannerOutputDailyPlanPoi, 0, len(schedule.POIs)),
		}

		for _, poi := range schedule.POIs {
			poiID, err := uuid.Parse(poi.ID)
			if err != nil {
				logger.Errorw("failed to parse POI ID", "err", err)
				return nil, fmt.Errorf("failed to parse POI ID: %v", err)
			}

			dailyItinerary.POIs = append(dailyItinerary.POIs, &do.TripPlannerOutputDailyPlanPoi{
				Id:    poiID,
				Notes: poi.Notes,
			})
		}

		output.DailyItineraries = append(output.DailyItineraries, dailyItinerary)
	}

	logger.Debugw("created trip planner output", "output", output)
	return output, nil
}

// Execute implements the main logic of PlanAgent
func (a *tripPlanner) Execute(ctx context.Context, inputAny any) (any, error) {
	input, ok := inputAny.(*do.TripPlannerInput)
	if !ok {
		a.logger.Errorw("invalid input type", "type", fmt.Sprintf("%T", inputAny))
		return nil, fmt.Errorf("invalid input type: %T", inputAny)
	}

	// get completion tool
	completionTool, err := a.toolHub.GetTool(cnst.ToolCompletion)
	if err != nil {
		a.logger.Errorw("failed to get completion tool", "err", err)
		return nil, err
	}

	// refine trip plan
	itineraries, err := getItineraries(ctx, completionTool, a.GetPrompt(), input, input.POIs, a.toolHub, a.GetToolNames(), a.logger)
	if err != nil {
		a.logger.Errorw("failed to get itineraries", "err", err)
		return nil, err
	}

	return itineraries, nil
}
