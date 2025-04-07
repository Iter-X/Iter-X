package tools

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"

	"github.com/iter-x/iter-x/internal/biz/ai/core"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/conf"
)

func NewCompletion(cfg *conf.Agent_ToolConfig) core.Tool {
	return &completionImpl{
		BaseTool: core.NewBaseTool(cfg.GetName().String(), cfg.GetDescription()),
		cli: openai.NewClient(
			option.WithAPIKey(cfg.GetApiKey()),
			option.WithBaseURL(cfg.GetBaseUrl()),
		),
		model: cfg.GetModel(),
	}
}

type completionImpl struct {
	*core.BaseTool
	cli   openai.Client
	model string
}

func (l completionImpl) Execute(ctx context.Context, inputAny any) (any, error) {
	var (
		input *do.ToolCompletionInput
		ok    bool
	)

	if input, ok = inputAny.(*do.ToolCompletionInput); !ok {
		return nil, fmt.Errorf("invalid input type: %T", input)
	}

	resp, err := l.cli.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: convertMessages(input.Messages),
		Model:    l.model,
		Tools:    convertTools(input.Tools),
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Choices) == 0 {
		return nil, nil
	}

	return &do.ToolCompletionOutput{Content: resp.Choices[0].Message.Content}, nil
}

func convertMessages(messages []do.ToolCompletionInputMessage) []openai.ChatCompletionMessageParamUnion {
	result := make([]openai.ChatCompletionMessageParamUnion, 0, len(messages))
	for _, msg := range messages {
		switch msg.Role {
		case do.CompletionRoleUser:
			result = append(result, openai.UserMessage(msg.Content))
		case do.CompletionRoleAssistant:
			result = append(result, openai.AssistantMessage(msg.Content))
		case do.CompletionRoleSystem:
			result = append(result, openai.SystemMessage(msg.Content))
		}
	}
	return result
}

// Related docs:
//
// https://platform.openai.com/docs/guides/function-calling
//
// https://docs.anthropic.com/en/docs/build-with-claude/tool-use/overview
//
// https://help.aliyun.com/zh/model-studio/qwen-function-calling
func convertTools(tools []do.FunctionCallTool) []openai.ChatCompletionToolParam {
	if len(tools) == 0 {
		return nil
	}

	result := make([]openai.ChatCompletionToolParam, 0, len(tools))
	for _, tool := range tools {
		result = append(result, openai.ChatCompletionToolParam{
			Function: openai.FunctionDefinitionParam{
				Name:        tool.Name,
				Description: openai.String(tool.Description),
				Parameters: openai.FunctionParameters{
					"type":       "object",
					"properties": tool.Parameters.Properties,
					"required":   tool.Parameters.Required,
				},
				Strict: openai.Bool(true),
			},
		})
	}
	return result
}
