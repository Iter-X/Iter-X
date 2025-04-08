package ai

import (
	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/ai/agent"
	"github.com/iter-x/iter-x/internal/biz/ai/tool"
	"github.com/iter-x/iter-x/internal/conf"
)

// ProviderSet is wire providers.
var ProviderSet = wire.NewSet(
	NewAgentHub,
	NewToolHub,
)

func NewToolHub(cfg *conf.Agent, logger *zap.SugaredLogger) (*tool.Hub, error) {
	return tool.NewToolHub(cfg, logger)
}

func NewAgentHub(cfg *conf.Agent, toolHub *tool.Hub, logger *zap.SugaredLogger) (*agent.Hub, error) {
	return agent.NewHub(cfg, toolHub, logger)
}
