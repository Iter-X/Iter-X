package bo

// Agent represents an agent business object
type Agent struct {
	Name    string  `json:"name"`
	Enabled bool    `json:"enabled"`
	Prompt  *Prompt `json:"prompt,omitempty"`
	Tools   []*Tool `json:"tools"`
}

// Prompt represents a prompt business object
type Prompt struct {
	System string `json:"system"`
	User   string `json:"user"`
}

// Tool represents a tool business object
type Tool struct {
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	Enabled     bool                   `json:"enabled"`
	Description string                 `json:"description,omitempty"`
	BaseURL     string                 `json:"base_url,omitempty"`
	APIKey      string                 `json:"api_key,omitempty"`
	Model       string                 `json:"model,omitempty"`
	Function    map[string]interface{} `json:"function,omitempty"`
}

// Function represents a function specification for LLMUse tools
type Function struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  map[string]interface{} `json:"parameters"`
	Required    []string               `json:"required"`
}
