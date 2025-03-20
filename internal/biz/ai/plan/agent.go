package plan

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/iter-x/iter-x/internal/common/cnst"

	"github.com/ifuryst/lol"
	"github.com/iter-x/iter-x/internal/biz/ai/tool"
	"github.com/iter-x/iter-x/internal/biz/do"

	"text/template"

	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/biz/ai/core"
	"github.com/iter-x/iter-x/internal/biz/repository"
)

type (
	// agent is an intelligent travel planning agent
	agent struct {
		*core.BaseAgent
		toolHub *tool.Hub
		poiRepo repository.PointsOfInterestRepo
	}

	completionResp struct {
		DailyPlans []*dailyPlan `json:"daily_plans"`
	}
	dailyPlan struct {
		Day        int         `json:"day"`   // e.g. 1
		Title      string      `json:"title"` // e.g. "Cultural Tour"
		Activities []*activity `json:"activities"`
	}
	activity struct {
		Time          string `json:"time"`     // e.g. "09:00"
		Name          string `json:"name"`     // e.g. "The Palace of Versailles"
		DurationInSec uint64 `json:"duration"` // e.g. 10800 (3 hours)
		Notes         string `json:"notes"`    // e.g. "Visit the world-famous palace and gardens"
	}

	refinedResp struct {
		DailyPlans []*refinedDailyPlan `json:"daily_plans"`
	}
	refinedDailyPlan struct {
		Day        int                `json:"day"`
		Title      string             `json:"title"`
		Activities []*refinedActivity `json:"activities"`
	}
	refinedActivity struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Time          string `json:"time"`
		DurationInSec uint64 `json:"duration"`
		Notes         string `json:"notes"`
	}

	userPromptTpl struct {
		Destination string
		StartDate   time.Time
		EndDate     time.Time
		Duration    int
		Preferences string
		Budget      float64
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

	// Generate prompt
	prompt := a.GetPrompt()
	systemPrompt := prompt.GetSystemPrompt()
	userPrompt := prompt.GetUserPrompt()
	tmplData := userPromptTpl{
		Destination: input.Destination,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		Duration:    input.Duration,
		Preferences: input.Preferences,
		Budget:      input.Budget,
	}
	tmpl, err := template.New("user_prompt").Parse(userPrompt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user prompt template: %v", err)
	}

	var userPromptBuf bytes.Buffer
	if err := tmpl.Execute(&userPromptBuf, tmplData); err != nil {
		return nil, fmt.Errorf("failed to execute user prompt template: %v", err)
	}

	completionInput := &do.ToolCompletionInput{
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

	var rawTrip completionResp
	if err := json.Unmarshal([]byte(completionOutput.Content), &rawTrip); err != nil {
		return nil, fmt.Errorf("failed to unmarshal completion output: %v", err)
	}

	// Extract all activity names from the raw trip
	activityNames := make([]string, 0, len(rawTrip.DailyPlans)*3)
	for _, dp := range rawTrip.DailyPlans {
		for _, act := range dp.Activities {
			activityNames = append(activityNames, act.Name)
		}
	}
	activityNames = lol.UniqSlice(activityNames)

	// Search for all activities in the database
	pois, err := a.poiRepo.SearchPointsOfInterestByNamesFromES(ctx, activityNames)
	if err != nil {
		return nil, fmt.Errorf("failed to search points of interest: %v", err)
	}

	// Create a map of name to POI for quick lookup
	poiMap := make(map[string]*do.PointsOfInterest)
	for _, poi := range pois {
		poiMap[poi.Name] = poi
	}

	// Generate refine prompt with POI information
	refinePrompt := a.GetPrompt().GetRefinePrompt()
	type refinePromptData struct {
		Activities []struct {
			ID   string
			Name string
		}
	}
	refineData := refinePromptData{
		Activities: make([]struct {
			ID   string
			Name string
		}, 0, len(pois)),
	}

	for _, poi := range pois {
		refineData.Activities = append(refineData.Activities, briefPOI{
			ID:   poi.ID.String(),
			Name: poi.Name,
		})
	}

	tmpl, err = template.New("refine_prompt").Parse(refinePrompt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse refine prompt template: %v", err)
	}

	var refinePromptBuf bytes.Buffer
	if err := tmpl.Execute(&refinePromptBuf, refineData); err != nil {
		return nil, fmt.Errorf("failed to execute refine prompt template: %v", err)
	}

	// Get refined plan with POI IDs
	completionInput = &do.ToolCompletionInput{
		Messages: []do.ToolCompletionInputMessage{
			{
				Role:    do.CompletionRoleSystem,
				Content: systemPrompt,
			},
			{
				Role:    do.CompletionRoleUser,
				Content: userPromptBuf.String(),
			},
			{
				Role:    do.CompletionRoleAssistant,
				Content: completionOutput.Content,
			},
			{
				Role:    do.CompletionRoleUser,
				Content: refinePromptBuf.String(),
			},
		},
	}
	resp, err = completionTool.Execute(ctx, completionInput)
	if err != nil {
		return nil, err
	}

	completionOutput, ok = resp.(*do.ToolCompletionOutput)
	if !ok {
		return nil, fmt.Errorf("invalid completion response type: %T", resp)
	}

	var refinedTrip refinedResp
	if err := json.Unmarshal([]byte(completionOutput.Content), &refinedTrip); err != nil {
		return nil, fmt.Errorf("failed to unmarshal refined completion output: %v", err)
	}

	// Convert to PlanAgentOutput
	result := &do.PlanAgentOutput{
		DailyPlans: make([]*do.DailyPlan, 0, len(refinedTrip.DailyPlans)),
	}

	for _, dp := range refinedTrip.DailyPlans {
		dailyPlan := &do.DailyPlan{
			Day:        dp.Day,
			Date:       input.StartDate.AddDate(0, 0, dp.Day-1),
			Title:      dp.Title,
			Activities: make([]*do.Activity, 0, len(dp.Activities)),
		}

		for _, act := range dp.Activities {
			// Parse time string to time.Time
			timeStr := fmt.Sprintf("%s %s", dailyPlan.Date.Format("2006-01-02"), act.Time)
			parsedTime, err := time.Parse("2006-01-02 15:04", timeStr)
			if err != nil {
				return nil, fmt.Errorf("failed to parse time %s: %v", act.Time, err)
			}

			// Parse POI ID
			var poiID uuid.UUID
			if act.ID != "" {
				poiID, err = uuid.Parse(act.ID)
				if err != nil {
					return nil, fmt.Errorf("failed to parse POI ID %s: %v", act.ID, err)
				}
			}

			duration := time.Duration(act.DurationInSec) * time.Second
			activity := &do.Activity{
				Id:       poiID,
				Time:     parsedTime,
				Name:     act.Name,
				Duration: duration,
				Notes:    act.Notes,
			}
			dailyPlan.Activities = append(dailyPlan.Activities, activity)
		}
		result.DailyPlans = append(result.DailyPlans, dailyPlan)
	}

	return result, nil
}
