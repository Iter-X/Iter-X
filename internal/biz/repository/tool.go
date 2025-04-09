package repository

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type Tool[T *ent.AgentTool, R *do.Tool] interface {
	BaseRepo[T, R]

	// GetTool returns a tool by name
	GetTool(ctx context.Context, name string) (*do.Tool, error)
	// ListTools returns all tools
	ListTools(ctx context.Context) ([]*do.Tool, error)
	// CreateTool creates a new tool
	CreateTool(ctx context.Context, tool *do.Tool) error
	// UpdateTool updates an existing tool
	UpdateTool(ctx context.Context, tool *do.Tool) error
	// DeleteTool deletes a tool
	DeleteTool(ctx context.Context, name string) error
}

type ToolRepo = Tool[*ent.AgentTool, *do.Tool]
