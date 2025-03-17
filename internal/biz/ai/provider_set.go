package ai

import (
	"github.com/google/wire"
	"github.com/iter-x/iter-x/internal/biz/ai/agent"
	"github.com/iter-x/iter-x/internal/biz/ai/tool"
)

// ProviderSet is wire providers.
var ProviderSet = wire.NewSet(agent.NewHub, tool.NewToolHub)
