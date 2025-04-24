package agent

import (
	"context"
	"fmt"
	"sync"

	"github.com/iter-x/iter-x/internal/biz/repository"

	"github.com/iter-x/iter-x/internal/biz/ai/core"
	"github.com/iter-x/iter-x/internal/biz/ai/planner/city"
	"github.com/iter-x/iter-x/internal/biz/ai/planner/trip"
	"github.com/iter-x/iter-x/internal/biz/ai/tool"
	"github.com/iter-x/iter-x/internal/common/cnst"
	"go.uber.org/zap"
)

// Hub is the central manager for all agents
type Hub struct {
	agents    map[string]core.Agent
	toolHub   *tool.Hub
	agentRepo repository.AgentRepo
	mu        sync.RWMutex
	logger    *zap.SugaredLogger
}

// NewHub initializes the AgentHub with configured agents
func NewHub(agentRepo repository.AgentRepo, toolHub *tool.Hub, logger *zap.SugaredLogger) (*Hub, error) {
	hub := &Hub{
		agents:    make(map[string]core.Agent),
		toolHub:   toolHub,
		agentRepo: agentRepo,
		logger:    logger.Named("agent.hub"),
	}

	// Get all agents from database
	agents, err := agentRepo.ListAgents(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to list agents: %v", err)
	}

	// Register all agents
	for _, agentDo := range agents {
		if !agentDo.Enabled {
			hub.logger.Infow("agent disabled, skipping", "name", agentDo.Name)
			continue
		}

		// Create prompt
		prompt := NewPrompt(agentDo.Prompt.System, agentDo.Prompt.User, agentDo.Prompt.Version)

		// Get tool names
		toolNames := make([]string, 0)
		for _, t := range agentDo.Tools {
			if t.Type == "LLMUse" {
				toolNames = append(toolNames, t.Name)
			}
		}

		// Create agent
		agent, err := createAgent(agentDo.Name, hub.toolHub, prompt, toolNames, logger)
		if err != nil {
			hub.logger.Errorw("failed to create agent", "name", agentDo.Name, "err", err)
			return nil, fmt.Errorf("failed to create agent %s: %v", agentDo.Name, err)
		}

		// Register agent
		if err := hub.RegisterAgent(agent); err != nil {
			hub.logger.Errorw("failed to register agent", "name", agentDo.Name, "err", err)
			return nil, fmt.Errorf("failed to register agent %s: %v", agentDo.Name, err)
		}
	}

	return hub, nil
}

// createAgent creates an agent based on configuration
func createAgent(name string, toolHub *tool.Hub, prompt core.Prompt, toolNames []string, logger *zap.SugaredLogger) (core.Agent, error) {
	switch name {
	case cnst.AgentTripPlanner:
		return trip.NewTripPlanner(name, toolHub, prompt, toolNames, logger), nil
	case cnst.AgentCityPlanner:
		return city.NewCityPlanner(name, toolHub, prompt, toolNames, logger), nil
	default:
		return nil, fmt.Errorf("unsupported agent: %s", name)
	}
}

// RegisterAgent registers a new agent to the hub
func (h *Hub) RegisterAgent(agent core.Agent) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	name := agent.Name()
	if _, exists := h.agents[name]; exists {
		h.logger.Errorw("agent already registered", "name", name)
		return fmt.Errorf("agent %s already registered", name)
	}

	h.agents[name] = agent
	h.logger.Infow("agent registered successfully", "name", name)
	return nil
}

// ListAgents returns all registered agents
func (h *Hub) ListAgents() []core.Agent {
	h.mu.RLock()
	defer h.mu.RUnlock()

	agents := make([]core.Agent, 0, len(h.agents))
	for _, agent := range h.agents {
		agents = append(agents, agent)
	}
	return agents
}

// GetAgent returns an agent by name
func (h *Hub) GetAgent(name string) (core.Agent, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	agent, exists := h.agents[name]
	if !exists {
		h.logger.Errorw("agent not found", "name", name)
		return nil, fmt.Errorf("agent %s not found", name)
	}

	return agent, nil
}
