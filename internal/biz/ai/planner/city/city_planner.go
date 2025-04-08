package city

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"text/template"
	"time"

	"github.com/iter-x/iter-x/internal/biz/ai/core"
	"github.com/iter-x/iter-x/internal/biz/ai/tool"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/common/cnst"
	"go.uber.org/zap"
)

type (
	// cityAgent is an intelligent travel planning agent
	cityAgent struct {
		*core.BaseAgent
		toolHub *tool.Hub
		logger  *zap.SugaredLogger
	}

	r1UserPromptTpl struct {
		Destination string
		StartDate   time.Time
		EndDate     time.Time
		Duration    int
		Preferences string
		Budget      float64
	}
)

// NewCityPlanner creates a new city planner agent
func NewCityPlanner(name string, toolHub *tool.Hub, prompt core.Prompt, toolNames []string, logger *zap.SugaredLogger) core.Agent {
	agent := &cityAgent{
		BaseAgent: core.NewBaseAgent(name, prompt, toolNames),
		toolHub:   toolHub,
		logger:    logger.Named("agent.city_planner"),
	}
	return agent
}

func getPrompt(prompt core.Prompt, input *do.CityPlannerInput, toolHub *tool.Hub,
	toolNames []string) (*do.ToolCompletionInput, error) {
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
	tmpl, err := template.New("user_prompt").Parse(userPrompt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user prompt template: %v", err)
	}

	var userPromptBuf bytes.Buffer
	if err := tmpl.Execute(&userPromptBuf, tmplData); err != nil {
		return nil, fmt.Errorf("failed to execute user prompt template: %v", err)
	}

	// Get all tools from the agent's configuration
	tools := make([]do.FunctionCallTool, 0, len(toolNames))
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

func getPotentialCities(ctx context.Context, completionTool core.Tool, prompt core.Prompt, input *do.CityPlannerInput,
	toolHub *tool.Hub, toolNames []string, logger *zap.SugaredLogger) ([][]string, error) {
	completionInput, err := getPrompt(prompt, input, toolHub, toolNames)
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
	completionOutput, ok := resp.(*do.ToolCompletionOutput[string])
	if !ok {
		logger.Errorw("invalid completion response type", "type", fmt.Sprintf("%T", resp))
		return nil, fmt.Errorf("invalid completion response type: %T", resp)
	}

	var cities [][]string
	if err := json.Unmarshal([]byte(completionOutput.Content), &cities); err != nil {
		logger.Errorw("failed to unmarshal cities", "err", err)
		return nil, err
	}
	logger.Debugw("got potential cities", "cities", cities)
	return cities, nil
}

// Execute implements the main logic of PlanAgent
func (a *cityAgent) Execute(ctx context.Context, inputAny any) (any, error) {
	input, ok := inputAny.(*do.CityPlannerInput)
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

	// get potential cities
	potentialCities, err := getPotentialCities(ctx, completionTool, a.GetPrompt(), input, a.toolHub, a.GetToolNames(), a.logger)
	if err != nil {
		a.logger.Errorw("failed to get potential cities", "err", err)
		return nil, err
	}

	// Convert to TripPlannerOutput
	result := do.CityPlannerOutput(potentialCities)
	a.logger.Debugw("created city planner output", "result", result)
	return &result, nil
}
