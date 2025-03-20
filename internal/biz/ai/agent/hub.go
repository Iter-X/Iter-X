package agent

import (
	"fmt"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"sync"

	"github.com/iter-x/iter-x/internal/biz/ai/core"
	"github.com/iter-x/iter-x/internal/biz/ai/plan"
	"github.com/iter-x/iter-x/internal/biz/ai/tool"
	"github.com/iter-x/iter-x/internal/conf"
)

// Hub is the central manager for all agents
type Hub struct {
	agents  map[string]core.Agent
	toolHub *tool.Hub
	mu      sync.RWMutex
}

// NewHub initializes the AgentHub with configured agents
func NewHub(cfg *conf.Agent, toolHub *tool.Hub, poiRepo repository.PointsOfInterestRepo) (*Hub, error) {
	hub := &Hub{
		agents:  make(map[string]core.Agent),
		toolHub: toolHub,
	}

	// Register all configured agents
	for _, agentCfg := range cfg.GetAgents() {
		if !agentCfg.GetEnabled() {
			continue
		}

		// Create prompt
		prompt := NewPrompt(
			agentCfg.GetPrompt().GetSystemPrompt(),
			agentCfg.GetPrompt().GetUserPrompt(),
			agentCfg.GetPrompt().GetRefinePrompt(),
			agentCfg.GetVersion(),
		)

		// Create agent
		agent, err := createAgent(agentCfg, hub.toolHub, prompt, poiRepo)
		if err != nil {
			return nil, fmt.Errorf("failed to create agent %s: %v", agentCfg.GetName(), err)
		}

		// Register agent
		if err := hub.RegisterAgent(agent); err != nil {
			return nil, fmt.Errorf("failed to register agent %s: %v", agentCfg.GetName(), err)
		}
	}

	return hub, nil
}

// createAgent creates an agent based on configuration
func createAgent(cfg *conf.Agent_AgentConfig, toolHub *tool.Hub, prompt core.Prompt, poiRepo repository.PointsOfInterestRepo) (core.Agent, error) {
	switch cfg.GetName() {
	case conf.Agent_PlanAgent:
		return plan.NewAgent(cfg.GetName().String(), toolHub, prompt, poiRepo), nil
	default:
		return nil, fmt.Errorf("unsupported agent: %s", cfg.GetName())
	}
}

// RegisterAgent registers a new agent to the hub
func (h *Hub) RegisterAgent(agent core.Agent) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	name := agent.Name()
	if _, exists := h.agents[name]; exists {
		return fmt.Errorf("agent %s already registered", name)
	}

	h.agents[name] = agent
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
		return nil, fmt.Errorf("agent %s not found", name)
	}

	return agent, nil
}
