package core

import (
	"context"
	"fmt"

	"github.com/iter-x/iter-x/internal/biz/do"
	"go.uber.org/zap"
)

// BaseTool provides a basic implementation of the Tool interface
type BaseTool struct {
	name   string
	desc   string
	Logger *zap.SugaredLogger
}

// NewBaseTool creates a new BaseTool instance
func NewBaseTool(name, desc string, logger *zap.SugaredLogger) *BaseTool {
	return &BaseTool{
		name:   name,
		desc:   desc,
		Logger: logger,
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

// GetLogger returns the logger instance
func (a *BaseTool) GetLogger() *zap.SugaredLogger {
	return a.Logger
}
