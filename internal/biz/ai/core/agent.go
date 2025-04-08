package core

import (
	"context"
	"fmt"
)

// BaseAgent provides a basic implementation of the Agent interface
type BaseAgent struct {
	name      string
	prompt    Prompt
	toolNames []string
}

// NewBaseAgent creates a new BaseAgent instance
func NewBaseAgent(name string, prompt Prompt, toolNames []string) *BaseAgent {
	return &BaseAgent{
		name:      name,
		prompt:    prompt,
		toolNames: toolNames,
	}
}

// Name returns the name of the agent
func (a *BaseAgent) Name() string {
	return a.name
}

// GetPrompt returns the prompt used by the agent
func (a *BaseAgent) GetPrompt() Prompt {
	return a.prompt
}

// GetToolNames returns the names of tools used by the agent
func (a *BaseAgent) GetToolNames() []string {
	return a.toolNames
}

// Execute provides a basic implementation that can be overridden by subclasses
func (a *BaseAgent) Execute(context.Context, any) (any, error) {
	return nil, fmt.Errorf("not implemented")
}
