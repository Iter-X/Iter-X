package plan

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/iter-x/iter-x/internal/common/cnst"
	"time"

	"github.com/ifuryst/lol"
	"github.com/iter-x/iter-x/internal/biz/ai/tool"
	"github.com/iter-x/iter-x/internal/biz/do"

	"text/template"

	"github.com/iter-x/iter-x/internal/biz/ai/core"
)

type (
	// agent is an intelligent travel planning agent
	agent struct {
		*core.BaseAgent
		toolHub *tool.Hub
	}

	completionResp struct {
		DailyPlans []*dailyPlan
	}
	dailyPlan struct {
		Day        int
		Date       time.Time
		Title      string
		Activities []*activity
	}
	activity struct {
		Time     time.Time
		Name     string
		Duration time.Duration
		Notes    string
	}

	userPromptTpl struct {
		Destination string
		StartDate   time.Time
		EndDate     time.Time
		Duration    int
		Preferences string
		Budget      float64
	}
)

// NewAgent creates a new PlanAgent
func NewAgent(name string, toolHub *tool.Hub, prompt core.Prompt) core.Agent {
	return &agent{
		BaseAgent: core.NewBaseAgent(name, prompt),
		toolHub:   toolHub,
	}
}

// Execute implements the main logic of PlanAgent
func (a *agent) Execute(ctx context.Context, inputAny any) (any, error) {
	// Implement the specific travel planning logic here
	// 1. Parse input parameters
	// 2. Use the toolset for planning
	// 3. Generate travel plan
	// 4. Return the result

	input, ok := inputAny.(*do.PlanAgentInput)
	if !ok {
		return nil, fmt.Errorf("invalid input type: %T", inputAny)
	}

	// get completion tool
	completionTool, err := a.toolHub.GetTool(cnst.ToolCompletion)
	if err != nil {
		return nil, err
	}

	// 处理prompt模板
	prompt := a.GetPrompt()
	systemPrompt := prompt.GetSystemPrompt()
	userPrompt := prompt.GetUserPrompt()

	// 创建模板数据
	tmplData := userPromptTpl{
		Destination: input.Destination,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		Duration:    input.Duration,
		Preferences: input.Preferences,
		Budget:      input.Budget,
	}

	// 处理用户prompt模板
	tmpl, err := template.New("user_prompt").Parse(userPrompt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user prompt template: %v", err)
	}

	var userPromptBuf bytes.Buffer
	if err := tmpl.Execute(&userPromptBuf, tmplData); err != nil {
		return nil, fmt.Errorf("failed to execute user prompt template: %v", err)
	}

	// 组装completion input
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
	// TODO: Implement activity search

	// Readjust the activities based on the search results

	// Return the final trip plan
	return nil, nil
}
