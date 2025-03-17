package bo

type (
	ToolCompletionInputMessage struct {
		Role    CompletionRole
		Content string
	}
	ToolCompletionInput struct {
		Messages []ToolCompletionInputMessage
	}
	ToolCompletionOutput struct {
		Content string
	}
)

type CompletionRole string

const (
	CompletionRoleUser      CompletionRole = "user"
	CompletionRoleAssistant CompletionRole = "assistant"
	CompletionRoleSystem    CompletionRole = "system"
)
