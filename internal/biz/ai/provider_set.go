package ai

import (
	"github.com/google/wire"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/ai/agent"
	"github.com/iter-x/iter-x/internal/biz/ai/tool"
)

// ProviderSet is wire providers.
var ProviderSet = wire.NewSet(
	NewAgentHub,
	NewToolHub,
)

func NewToolHub(toolRepo repository.ToolRepo, logger *zap.SugaredLogger) (*tool.Hub, error) {
	return tool.NewToolHub(toolRepo, logger)
}

func NewAgentHub(toolHub *tool.Hub, agentRepo repository.AgentRepo, logger *zap.SugaredLogger) (*agent.Hub, error) {
	return agent.NewHub(agentRepo, toolHub, logger)
}
