package core

import (
	"context"
	"fmt"
	"github.com/iter-x/iter-x/internal/biz/do"
)

// BaseTool provides a basic implementation of the Tool interface
type BaseTool struct {
	name string
	desc string
}

// NewBaseTool creates a new BaseTool instance
func NewBaseTool(name, desc string) *BaseTool {
	return &BaseTool{
		name: name,
		desc: desc,
	}
}

// Name returns the name of the tool
func (a *BaseTool) Name() string {
	return a.name
}

// Description returns the description of the tool
func (a *BaseTool) Description() string {
	return a.desc
}

// Execute provides a basic implementation that can be overridden by subclasses
func (a *BaseTool) Execute(context.Context, any) (any, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetDefinition returns the definition of the tool, only for function call tools
func (a *BaseTool) GetDefinition() (*do.FunctionCallTool, error) {
	return nil, fmt.Errorf("not implemented")
}
