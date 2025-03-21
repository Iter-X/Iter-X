package tool

import (
	"fmt"
	"sync"

	"github.com/iter-x/iter-x/internal/biz/ai/core"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/data/impl/tools"
)

// Hub is the central manager for all tools
type Hub struct {
	tools map[string]core.Tool
	mu    sync.RWMutex
}

// NewToolHub creates a new Hub instance
func NewToolHub(cfg *conf.Agent) (*Hub, error) {
	hub := &Hub{
		tools: make(map[string]core.Tool),
	}

	for _, toolCfg := range cfg.GetTools() {
		t, err := createTool(toolCfg)
		if err != nil {
			return nil, fmt.Errorf("failed to create tool %s: %v", toolCfg.GetName(), err)
		}

		if err := hub.RegisterTool(t); err != nil {
			return nil, fmt.Errorf("failed to register tool %s: %v", toolCfg.GetName(), err)
		}
	}

	return hub, nil
}

// RegisterTool registers a new tool to the hub
func (h *Hub) RegisterTool(tool core.Tool) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	name := tool.Name()
	if _, exists := h.tools[name]; exists {
		return fmt.Errorf("tool %s already registered", name)
	}

	h.tools[name] = tool
	return nil
}

// createTool creates a tool based on configuration
func createTool(cfg *conf.Agent_ToolConfig) (core.Tool, error) {
	switch cfg.GetName() {
	case conf.Agent_Completion:
		return tools.NewCompletion(cfg), nil
	default:
		return nil, fmt.Errorf("unknown tool: %s", cfg.GetName())
	}
}

// GetTool returns a tool by name
func (h *Hub) GetTool(name string) (core.Tool, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	tool, exists := h.tools[name]
	if !exists {
		return nil, fmt.Errorf("tool %s not found", name)
	}

	return tool, nil
}
