package tool

import (
	"context"
	"fmt"
	"sync"

	"github.com/iter-x/iter-x/internal/biz/ai/core"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/common/cnst"
	"github.com/iter-x/iter-x/internal/data/impl/tools"
	"go.uber.org/zap"
)

// Hub is the central manager for all tools
type Hub struct {
	tools    map[string]core.Tool
	toolRepo repository.ToolRepo
	mu       sync.RWMutex
	logger   *zap.SugaredLogger
}

// NewToolHub creates a new Hub instance
func NewToolHub(toolRepo repository.ToolRepo, logger *zap.SugaredLogger) (*Hub, error) {
	hub := &Hub{
		tools:    make(map[string]core.Tool),
		toolRepo: toolRepo,
		logger:   logger.Named("tool.hub"),
	}

	// Get all tools from database
	allTool, err := toolRepo.ListTools(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to list allTool: %v", err)
	}

	// Register all tools
	for _, toolDo := range allTool {
		if !toolDo.Enabled {
			hub.logger.Infow("tool disabled, skipping", "name", toolDo.Name)
			continue
		}

		// Create tool
		t, err := createTool(toolDo, logger)
		if err != nil {
			hub.logger.Errorw("failed to create tool", "name", toolDo.Name, "err", err)
			return nil, fmt.Errorf("failed to create tool %s: %v", toolDo.Name, err)
		}

		// Register tool
		if err := hub.RegisterTool(t); err != nil {
			hub.logger.Errorw("failed to register tool", "name", toolDo.Name, "err", err)
			return nil, fmt.Errorf("failed to register tool %s: %v", toolDo.Name, err)
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
func createTool(toolDo *do.Tool, logger *zap.SugaredLogger) (core.Tool, error) {
	switch toolDo.Type {
	case cnst.ToolTypeCodeUse:
		return tools.NewCompletion(toolDo, logger), nil
	case cnst.ToolTypeLLMUse:
		return tools.NewLLMUse(toolDo, logger), nil
	default:
		return nil, fmt.Errorf("unknown tool type: %s", toolDo.Type)
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
