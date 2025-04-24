package repository

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type Agent[T *ent.Agent, R *do.Agent] interface {
	BaseRepo[T, R]

	// GetAgent returns an agent by name
	GetAgent(ctx context.Context, name string) (*do.Agent, error)
	// ListAgents returns all agents
	ListAgents(ctx context.Context) ([]*do.Agent, error)
	// CreateAgent creates a new agent
	CreateAgent(ctx context.Context, agent *do.Agent) error
	// UpdateAgent updates an existing agent
	UpdateAgent(ctx context.Context, agent *do.Agent) error
	// DeleteAgent deletes an agent
	DeleteAgent(ctx context.Context, name string) error
}

type AgentRepo = Agent[*ent.Agent, *do.Agent]
