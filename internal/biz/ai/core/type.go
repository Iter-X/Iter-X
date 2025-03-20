package core

import (
	"context"
)

// Prompt defines the interface for a prompt
type Prompt interface {
	GetSystemPrompt() string
	GetUserPrompt() string
	GetVersion() string
	GetRefinePrompt() string
}

type (
	// Agent defines the basic interface for an agent
	Agent interface {
		// Name returns the name of the agent
		Name() string
		// GetPrompt returns the prompt used by the agent
		GetPrompt() Prompt
		Execute(context.Context, any) (any, error)
	}
)

type (
	// Tool defines the interface for a tool
	Tool interface {
		// Name returns the name of the tool
		Name() string
		// Description returns the description of the tool
		Description() string
		Execute(context.Context, any) (any, error)
	}
)
