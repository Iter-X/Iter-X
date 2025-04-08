package tool

import (
	"fmt"
	"sync"

	"github.com/iter-x/iter-x/internal/biz/ai/core"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/data/impl/tools"
	"go.uber.org/zap"
)

// Hub is the central manager for all tools
type Hub struct {
	tools  map[string]core.Tool
	mu     sync.RWMutex
	logger *zap.SugaredLogger
}

// NewToolHub creates a new Hub instance
func NewToolHub(cfg *conf.Agent, logger *zap.SugaredLogger) (*Hub, error) {
	hub := &Hub{
		tools:  make(map[string]core.Tool),
		logger: logger.Named("tool.hub"),
	}

	// Register code use tools
	for _, toolCfg := range cfg.GetTools() {
		t, err := createTool(toolCfg, logger)
		if err != nil {
			hub.logger.Errorw("failed to create tool", "name", toolCfg.GetName(), "err", err)
			return nil, fmt.Errorf("failed to create tool %s: %v", toolCfg.GetName(), err)
		}

		if err := hub.RegisterTool(t); err != nil {
			hub.logger.Errorw("failed to register tool", "name", toolCfg.GetName(), "err", err)
			return nil, fmt.Errorf("failed to register tool %s: %v", toolCfg.GetName(), err)
		}
	}

	// Register LLM use tools
	for _, toolCfg := range cfg.GetLlmUseTools() {
		t := tools.NewLLMUse(toolCfg, logger)
		if err := hub.RegisterTool(t); err != nil {
			hub.logger.Errorw("failed to register tool", "name", toolCfg.GetName(), "err", err)
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
		h.logger.Errorw("tool already registered", "name", name)
		return fmt.Errorf("tool %s already registered", name)
	}

	h.tools[name] = tool
	h.logger.Infow("tool registered successfully", "name", name)
	return nil
}

// createTool creates a tool based on configuration
func createTool(cfg *conf.Agent_ToolConfig, logger *zap.SugaredLogger) (core.Tool, error) {
	switch cfg.GetName() {
	case conf.Agent_Completion:
		return tools.NewCompletion(cfg, logger), nil
	case conf.Agent_InducingCreateTrip:
		return tools.NewLLMUse(cfg.GetLlmUseConfig(), logger), nil
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
		h.logger.Errorw("tool not found", "name", name)
		return nil, fmt.Errorf("tool %s not found", name)
	}

	return tool, nil
}
