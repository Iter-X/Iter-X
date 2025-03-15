package agent

import (
	"context"
	"fmt"
	"sync"

	"github.com/iter-x/iter-x/internal/biz/agent/plan"
	"github.com/iter-x/iter-x/internal/conf"

	"github.com/iter-x/iter-x/internal/biz/agent/core"
)

// Hub is the central manager for all agents
type Hub struct {
	agents map[string]core.Agent
	mu     sync.RWMutex
}

// NewHub initializes the AgentHub with configured agents
func NewHub(cfg *conf.Agent) (*Hub, error) {
	hub := &Hub{
		agents: make(map[string]core.Agent),
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
			agentCfg.GetVersion(),
		)

		// Create tools
		var tools []core.Tool
		for _, toolCfg := range agentCfg.GetTools() {
			if !toolCfg.GetEnabled() {
				continue
			}

			tool, err := createTool(toolCfg)
			if err != nil {
				return nil, fmt.Errorf("failed to create tool %s: %v", toolCfg.GetName(), err)
			}
			tools = append(tools, tool)
		}

		// Create agent
		agent, err := createAgent(agentCfg.GetName(), agentCfg.GetName(), tools, prompt)
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
func createAgent(name, description string, tools []core.Tool, prompt core.Prompt) (core.Agent, error) {
	switch name {
	case core.AgentNamePlan.String():
		return plan.NewAgent(name, description, tools, prompt), nil
	default:
		return nil, fmt.Errorf("unsupported agent type: %s", name)
	}
}

// createTool creates a tool based on configuration
func createTool(cfg *conf.Agent_ToolConfig) (core.Tool, error) {
	switch cfg.GetName() {
	case conf.Agent_Completion:
		return nil, nil
	default:
		return nil, fmt.Errorf("unknown tool type: %s", cfg.GetName())
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

// GetAgent retrieves an agent by its name
func (h *Hub) GetAgent(name string) (core.Agent, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	agent, exists := h.agents[name]
	if !exists {
		return nil, fmt.Errorf("agent %s not found", name)
	}

	return agent, nil
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

// ExecuteAgent executes a task using the specified agent
func (h *Hub) ExecuteAgent(ctx context.Context, name string, input interface{}) (interface{}, error) {
	agent, err := h.GetAgent(name)
	if err != nil {
		return nil, err
	}

	return agent.Execute(ctx, input)
}

// GetAgentTools returns all tools available to the specified agent
func (h *Hub) GetAgentTools(name string) ([]core.Tool, error) {
	agent, err := h.GetAgent(name)
	if err != nil {
		return nil, err
	}

	return agent.GetTools(), nil
}

// GetAgentPrompt returns the prompt used by the specified agent
func (h *Hub) GetAgentPrompt(name string) (core.Prompt, error) {
	agent, err := h.GetAgent(name)
	if err != nil {
		return nil, err
	}

	return agent.GetPrompt(), nil
}
