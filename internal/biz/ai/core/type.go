package core

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/do"
)

// Prompt defines the interface for a prompt
type Prompt interface {
	GetSystemPrompt() string
	GetUserPrompt() string
	GetSystemPromptByRound(round int) string
	GetUserPromptByRound(round int) string
	GetVersion() string
}

type (
	// Agent defines the basic interface for an agent
	Agent interface {
		// Name returns the name of the agent
		Name() string
		// GetPrompt returns the prompt used by the agent
		GetPrompt() Prompt
		// GetToolNames returns the names of tools used by the agent
		GetToolNames() []string
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
		// GetDefinition returns the definition of the tool, only for function call tools
		GetDefinition() (*do.FunctionCallTool, error)
	}
)
