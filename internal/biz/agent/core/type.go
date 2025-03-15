package core

import (
	"context"
)

// Tool defines the interface for a tool
type Tool interface {
	Name() string
	Description() string
	Execute(ctx context.Context, input interface{}) (interface{}, error)
}

// Prompt defines the interface for a prompt
type Prompt interface {
	GetSystemPrompt() string
	GetUserPrompt() string
	GetVersion() string
}

// Agent defines the basic interface for an agent
type Agent interface {
	// Name returns the name of the agent
	Name() string
	// Description returns the description of the agent
	Description() string
	// GetTools returns the list of tools available to the agent
	GetTools() []Tool
	// GetPrompt returns the prompt used by the agent
	GetPrompt() Prompt
	// Execute performs the main logic of the agent
	Execute(ctx context.Context, input interface{}) (interface{}, error)
}

type AgentName string

const (
	AgentNamePlan     AgentName = "PlanAgent"
	AgentNameDayOpt   AgentName = "DayOptAgent"
	AgentNameDestRecs AgentName = "DestRecsAgent"
)

func (a AgentName) String() string {
	return string(a)
}

type AgentTool string

const (
	AgentToolBrowser AgentName = "Browser"
)
