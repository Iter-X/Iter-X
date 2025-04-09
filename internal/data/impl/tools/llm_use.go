package tools

import (
	"context"
	"fmt"

	"github.com/iter-x/iter-x/internal/biz/ai/core"
	"github.com/iter-x/iter-x/internal/biz/do"
	"go.uber.org/zap"
)

// NewLLMUse creates a new LLMUse tool instance
func NewLLMUse(toolDo *do.Tool, logger *zap.SugaredLogger) core.Tool {
	return &llmUseImpl{
		BaseTool: core.NewBaseTool(toolDo.Name, toolDo.Type, logger.Named("tool.llm_use")),
		function: toolDo.Function,
	}
}

type llmUseImpl struct {
	*core.BaseTool
	function *do.FunctionConfig
}

func (l *llmUseImpl) Execute(_ context.Context, inputAny any) (any, error) {
	var (
		input *do.ToolCompletionInput
		ok    bool
	)

	if input, ok = inputAny.(*do.ToolCompletionInput); !ok {
		l.Logger.Errorw("invalid input type", "type", fmt.Sprintf("%T", input))
		return nil, fmt.Errorf("invalid input type: %T", input)
	}

	// Convert function config to FunctionCallTool
	tool := do.FunctionCallTool{
		Name:        l.function.Name,
		Description: l.function.Description,
		Parameters: do.FunctionCallToolParameters{
			Type:       l.function.Parameters.Type,
			Properties: convertProperties(l.function.Parameters.Properties),
			Required:   l.function.Parameters.Required,
		},
	}

	l.Logger.Debugw("adding tool to input", "tool", tool)

	// Add the tool to input.Tools
	input.Tools = append(input.Tools, tool)

	return input, nil
}

func convertProperties(props map[string]*do.FunctionParameterProperty) map[string]do.FunctionCallToolProperty {
	result := make(map[string]do.FunctionCallToolProperty)

	for key, prop := range props {
		result[key] = do.FunctionCallToolProperty{
			Type:        prop.Type,
			Description: prop.Description,
		}
	}

	return result
}

func (l *llmUseImpl) GetDefinition() (*do.FunctionCallTool, error) {
	tool := &do.FunctionCallTool{
		Name:        l.function.Name,
		Description: l.function.Description,
		Parameters: do.FunctionCallToolParameters{
			Type:       l.function.Parameters.Type,
			Properties: convertProperties(l.function.Parameters.Properties),
			Required:   l.function.Parameters.Required,
		},
	}
	l.Logger.Debugw("getting tool definition", "tool", tool)
	return tool, nil
}
