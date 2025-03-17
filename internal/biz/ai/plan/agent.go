package plan

import (
	"context"
	"github.com/iter-x/iter-x/internal/biz/ai/tool"

	"github.com/iter-x/iter-x/internal/biz/ai/core"
)

// agent is an intelligent travel planning agent
type agent struct {
	*core.BaseAgent
	toolHub *tool.Hub
}

// NewAgent creates a new PlanAgent
func NewAgent(name string, toolHub *tool.Hub, prompt core.Prompt) core.Agent {
	return &agent{
		BaseAgent: core.NewBaseAgent(name, prompt),
		toolHub:   toolHub,
	}
}

// Execute implements the main logic of PlanAgent
func (a *agent) Execute(context.Context, any) (any, error) {
	// Implement the specific travel planning logic here
	// 1. Parse input parameters
	// 2. Use the toolset for planning
	// 3. Generate travel plan
	// 4. Return the result

	return nil, nil
}
