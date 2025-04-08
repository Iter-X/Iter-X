package do

type (
	ToolCompletionInputMessage struct {
		Role    CompletionRole
		Content string
	}
	ToolCompletionInput struct {
		Messages []ToolCompletionInputMessage
		Tools    []FunctionCallTool
	}
	ToolCompletionOutput[T any] struct {
		Content T
	}

	FunctionCallTool struct {
		Name        string                     `json:"name"`
		Description string                     `json:"description"`
		Parameters  FunctionCallToolParameters `json:"parameters"`
	}
	FunctionCallToolParameters struct {
		Type       string                              `json:"type"`
		Properties map[string]FunctionCallToolProperty `json:"properties"`
		Required   []string                            `json:"required"`
	}
	FunctionCallToolProperty struct {
		Type        string `json:"type"`
		Description string `json:"description"`
	}

	ToolCallOutput struct {
		ID       string                 `json:"id"`
		Function ToolCallOutputFunction `json:"function"`
		Type     string                 `json:"type"`
	}
	ToolCallOutputFunction struct {
		Arguments string `json:"arguments"`
		Name      string `json:"name"`
	}
)

type CompletionRole string

const (
	CompletionRoleUser      CompletionRole = "user"
	CompletionRoleAssistant CompletionRole = "assistant"
	CompletionRoleSystem    CompletionRole = "system"
)
