package tools

import (
	"context"
	"fmt"

	"github.com/iter-x/iter-x/internal/biz/ai/core"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/conf"
)

// NewLLMUse creates a new LLMUse tool instance
func NewLLMUse(cfg *conf.Agent_LLMUseToolConfig) core.Tool {
	return &llmUseImpl{
		BaseTool: core.NewBaseTool(cfg.GetName(), cfg.GetType()),
		function: cfg.GetFunction(),
	}
}

type llmUseImpl struct {
	*core.BaseTool
	function *conf.Agent_FunctionConfig
}

func (l llmUseImpl) Execute(_ context.Context, inputAny any) (any, error) {
	var (
		input *do.ToolCompletionInput
		ok    bool
	)

	if input, ok = inputAny.(*do.ToolCompletionInput); !ok {
		return nil, fmt.Errorf("invalid input type: %T", input)
	}

	// Convert function config to FunctionCallTool
	tool := do.FunctionCallTool{
		Name:        l.function.GetName(),
		Description: l.function.GetDescription(),
		Parameters: do.FunctionCallToolParameters{
			Type:       l.function.GetParameters().GetType(),
			Properties: convertProperties(l.function.GetParameters().GetProperties()),
			Required:   l.function.GetParameters().GetRequired(),
		},
	}

	// Add the tool to input.Tools
	input.Tools = append(input.Tools, tool)

	return input, nil
}

func convertProperties(props map[string]*conf.Agent_FunctionParameters_Property) map[string]do.FunctionCallToolProperty {
	result := make(map[string]do.FunctionCallToolProperty)

	for key, prop := range props {
		result[key] = do.FunctionCallToolProperty{
			Type:        prop.GetType(),
			Description: prop.GetDescription(),
		}
	}

	return result
}
