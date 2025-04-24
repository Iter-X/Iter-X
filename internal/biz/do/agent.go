package do

import "time"

type (
	// Agent represents an agent domain object
	Agent struct {
		ID          int       `json:"id"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Name        string    `json:"name"`
		Enabled     bool      `json:"enabled"`
		Description string    `json:"description"`
		Prompt      *Prompt   `json:"prompt,omitempty"`
		Tools       []*Tool   `json:"tools"`
	}

	// Prompt represents a prompt domain object
	Prompt struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		System    string    `json:"system"`
		User      string    `json:"user"`
		Version   string    `json:"version"`
	}

	// Tool represents a tool domain object
	Tool struct {
		ID          int             `json:"id"`
		CreatedAt   time.Time       `json:"created_at"`
		UpdatedAt   time.Time       `json:"updated_at"`
		Name        string          `json:"name"`
		Type        string          `json:"type"`
		Enabled     bool            `json:"enabled"`
		Description string          `json:"description,omitempty"`
		BaseURL     string          `json:"base_url,omitempty"`
		APIKey      string          `json:"api_key,omitempty"`
		Model       string          `json:"model,omitempty"`
		Function    *FunctionConfig `json:"function,omitempty"`
	}

	// FunctionConfig represents a function configuration
	FunctionConfig struct {
		Name        string             `json:"name"`
		Description string             `json:"description"`
		Parameters  FunctionParameters `json:"parameters"`
	}

	// FunctionParameters represents function parameters
	FunctionParameters struct {
		Type       string                                `json:"type"`
		Properties map[string]*FunctionParameterProperty `json:"properties"`
		Required   []string                              `json:"required"`
	}

	// FunctionParameterProperty represents a function parameter property
	FunctionParameterProperty struct {
		Type        string `json:"type"`
		Description string `json:"description"`
	}

	// ToolCompletionInputMessage represents a message in tool completion input
	ToolCompletionInputMessage struct {
		Role    CompletionRole `json:"role"`
		Content string         `json:"content"`
	}

	// ToolCompletionInput represents input for tool completion
	ToolCompletionInput struct {
		Messages []ToolCompletionInputMessage `json:"messages"`
		Tools    []FunctionCallTool           `json:"tools"`
	}

	// ToolCompletionOutput represents output from tool completion
	ToolCompletionOutput[T any] struct {
		Content T `json:"content"`
	}

	// FunctionCallTool represents a function call tool
	FunctionCallTool struct {
		Name        string                     `json:"name"`
		Description string                     `json:"description"`
		Parameters  FunctionCallToolParameters `json:"parameters"`
	}

	// FunctionCallToolParameters represents parameters for a function call tool
	FunctionCallToolParameters struct {
		Type       string                              `json:"type"`
		Properties map[string]FunctionCallToolProperty `json:"properties"`
		Required   []string                            `json:"required"`
	}

	// FunctionCallToolProperty represents a property in function call tool parameters
	FunctionCallToolProperty struct {
		Type        string `json:"type"`
		Description string `json:"description"`
	}

	// ToolCallOutput represents output from a tool call
	ToolCallOutput struct {
		ID       string                 `json:"id"`
		Function ToolCallOutputFunction `json:"function"`
		Type     string                 `json:"type"`
	}

	// ToolCallOutputFunction represents a function in tool call output
	ToolCallOutputFunction struct {
		Arguments string `json:"arguments"`
		Name      string `json:"name"`
	}
)

// CompletionRole represents the role in a completion
type CompletionRole string

const (
	CompletionRoleUser      CompletionRole = "user"
	CompletionRoleAssistant CompletionRole = "assistant"
	CompletionRoleSystem    CompletionRole = "system"
)
