package tools

import (
	"context"
	"fmt"
	"github.com/iter-x/iter-x/internal/biz/ai/core"
	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
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
	cli   *openai.Client
	model string
}

func (l completionImpl) Execute(ctx context.Context, inputAny any) (any, error) {
	var (
		input  bo.ToolCompletionInput
		output bo.ToolCompletionOutput
		ok     bool
	)

	if input, ok = inputAny.(bo.ToolCompletionInput); !ok {
		return nil, fmt.Errorf("invalid input type: %T", input)
	}

	messages := make([]openai.ChatCompletionMessageParamUnion, 0, len(input.Messages))
	for _, msg := range input.Messages {
		switch msg.Role {
		case bo.CompletionRoleUser:
			messages = append(messages, openai.UserMessage(msg.Content))
		case bo.CompletionRoleAssistant:
			messages = append(messages, openai.AssistantMessage(msg.Content))
		case bo.CompletionRoleSystem:
			messages = append(messages, openai.SystemMessage(msg.Content))
		}
	}

	resp, err := l.cli.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F(messages),
		Model:    openai.F(l.model),
	})
	if err != nil {
		return output, err
	}

	if len(resp.Choices) == 0 {
		return output, nil
	}

	return bo.ToolCompletionOutput{Content: resp.Choices[0].Message.Content}, nil
}
