package core

import "context"

// BaseAgent provides a basic implementation of the Agent interface
type BaseAgent struct {
	name        string
	description string
	tools       []Tool
	prompt      Prompt
}

// NewBaseAgent creates a new BaseAgent instance
func NewBaseAgent(name, description string, tools []Tool, prompt Prompt) *BaseAgent {
	return &BaseAgent{
		name:        name,
		description: description,
		tools:       tools,
		prompt:      prompt,
	}
}

// Name returns the name of the agent
func (a *BaseAgent) Name() string {
	return a.name
}

// Description returns the description of the agent
func (a *BaseAgent) Description() string {
	return a.description
}

// GetTools returns the list of tools available to the agent
func (a *BaseAgent) GetTools() []Tool {
	return a.tools
}

// GetPrompt returns the prompt used by the agent
func (a *BaseAgent) GetPrompt() Prompt {
	return a.prompt
}

// Execute provides a basic implementation that can be overridden by subclasses
func (a *BaseAgent) Execute(_ context.Context, _ interface{}) (interface{}, error) {
	return nil, nil
}
