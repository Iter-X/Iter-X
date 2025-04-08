package agent

import (
	"fmt"
	"sync"

	"github.com/iter-x/iter-x/internal/biz/ai/core"
	"github.com/iter-x/iter-x/internal/biz/ai/planner/city"
	"github.com/iter-x/iter-x/internal/biz/ai/planner/trip"
	"github.com/iter-x/iter-x/internal/biz/ai/tool"
	"github.com/iter-x/iter-x/internal/conf"
	"go.uber.org/zap"
)

// Hub is the central manager for all agents
type Hub struct {
	agents  map[string]core.Agent
	toolHub *tool.Hub
	mu      sync.RWMutex
	logger  *zap.SugaredLogger
}

// NewHub initializes the AgentHub with configured agents
func NewHub(cfg *conf.Agent, toolHub *tool.Hub, logger *zap.SugaredLogger) (*Hub, error) {
	hub := &Hub{
		agents:  make(map[string]core.Agent),
		toolHub: toolHub,
		logger:  logger.Named("agent.hub"),
	}

	// Register all configured agents
	for _, agentCfg := range cfg.GetAgents() {
		if !agentCfg.GetEnabled() {
			hub.logger.Infow("agent disabled, skipping", "name", agentCfg.GetName())
			continue
		}

		// Create prompt
		rounds := make([]PromptRound, 0, len(agentCfg.GetPrompt().GetRounds()))
		for _, r := range agentCfg.GetPrompt().GetRounds() {
			rounds = append(rounds, PromptRound{
				System: r.GetSystem(),
				User:   r.GetUser(),
			})
		}
		prompt := NewPrompt(rounds, agentCfg.GetVersion())

		// Create agent
		agent, err := createAgent(agentCfg, hub.toolHub, prompt, logger)
		if err != nil {
			hub.logger.Errorw("failed to create agent", "name", agentCfg.GetName(), "err", err)
			return nil, fmt.Errorf("failed to create agent %s: %v", agentCfg.GetName(), err)
		}

		// Register agent
		if err := hub.RegisterAgent(agent); err != nil {
			hub.logger.Errorw("failed to register agent", "name", agentCfg.GetName(), "err", err)
			return nil, fmt.Errorf("failed to register agent %s: %v", agentCfg.GetName(), err)
		}
	}

	return hub, nil
}

// createAgent creates an agent based on configuration
func createAgent(cfg *conf.Agent_AgentConfig, toolHub *tool.Hub, prompt core.Prompt, logger *zap.SugaredLogger) (core.Agent, error) {
	toolNames := make([]string, 0)
	for _, toolCfg := range cfg.GetTools() {
		if toolCfg.Type != conf.Agent_AgentToolConfig_LLMUse {
			continue
		}
		toolNames = append(toolNames, toolCfg.GetName().String())
	}

	switch cfg.GetName() {
	case conf.Agent_TripPlanner:
		return trip.NewTripPlanner(cfg.GetName().String(), toolHub, prompt, toolNames, logger), nil
	case conf.Agent_CityPlanner:
		return city.NewCityPlanner(cfg.GetName().String(), toolHub, prompt, toolNames, logger), nil
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
