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

	messages := make([]openai.ChatCompletionMessageParamUnion, 0, len(input.Messages))
	for _, msg := range input.Messages {
		switch msg.Role {
		case do.CompletionRoleUser:
			messages = append(messages, openai.UserMessage(msg.Content))
		case do.CompletionRoleAssistant:
			messages = append(messages, openai.AssistantMessage(msg.Content))
		case do.CompletionRoleSystem:
			messages = append(messages, openai.SystemMessage(msg.Content))
		}
	}

	resp, err := l.cli.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    l.model,
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Choices) == 0 {
		return nil, nil
	}

	return &do.ToolCompletionOutput{Content: resp.Choices[0].Message.Content}, nil
}
