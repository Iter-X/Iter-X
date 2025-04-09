package tools

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/ai/core"
	"github.com/iter-x/iter-x/internal/biz/do"
)

func NewCompletion(toolDo *do.Tool, logger *zap.SugaredLogger) core.Tool {
	logger = logger.Named("tool.completion")
	return &completionImpl{
		BaseTool: core.NewBaseTool(toolDo.Name, toolDo.Description, logger),
		cli: openai.NewClient(
			option.WithAPIKey(toolDo.APIKey),
			option.WithBaseURL(toolDo.BaseURL),
		),
		model: toolDo.Model,
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
		l.Logger.Errorw("invalid input type", "type", fmt.Sprintf("%T", input))
		return nil, fmt.Errorf("invalid input type: %T", input)
	}

	l.Logger.Debugw("executing completion", "model", l.model, "messages", input.Messages)
	resp, err := l.cli.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: convertMessages(input.Messages),
		Model:    l.model,
		Tools:    convertTools(input.Tools),
	})
	if err != nil {
		l.Logger.Errorw("failed to execute completion", "err", err)
		return nil, err
	}

	if len(resp.Choices) == 0 {
		l.Logger.Warnw("no choices returned from completion")
		return nil, nil
	}

	if resp.Choices[0].FinishReason == "function_call" || resp.Choices[0].FinishReason == "tool_calls" {
		res := make([]do.ToolCallOutput, 0, len(resp.Choices[0].Message.ToolCalls))
		for _, t := range resp.Choices[0].Message.ToolCalls {
			res = append(res, do.ToolCallOutput{
				ID: t.ID,
				Function: do.ToolCallOutputFunction{
					Arguments: t.Function.Arguments,
					Name:      t.Function.Name,
				},
				Type: string(t.Type),
			})
		}
		l.Logger.Debugw("completion returned tool calls", "tool_calls", res)
		return &do.ToolCompletionOutput[[]do.ToolCallOutput]{
			Content: res,
		}, nil
	}

	l.Logger.Debugw("completion returned content", "content", resp.Choices[0].Message.Content)
	return &do.ToolCompletionOutput[string]{Content: resp.Choices[0].Message.Content}, nil
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
