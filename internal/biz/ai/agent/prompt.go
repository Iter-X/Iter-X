package agent

import (
	"fmt"
	"sync"

	"github.com/iter-x/iter-x/internal/biz/ai/core"
)

// PromptVersion represents a version of a prompt
type PromptVersion struct {
	Version      string
	SystemPrompt string
	UserPrompt   string
	RefinePrompt string
}

// PromptManager manages different versions of prompts
type PromptManager struct {
	prompts map[string]map[string]PromptVersion
	mu      sync.RWMutex
}

// NewPromptManager creates a new prompt manager
func NewPromptManager() *PromptManager {
	return &PromptManager{
		prompts: make(map[string]map[string]PromptVersion),
	}
}

// RegisterPrompt registers a new prompt version for an agent
func (pm *PromptManager) RegisterPrompt(agentName string, version PromptVersion) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if _, exists := pm.prompts[agentName]; !exists {
		pm.prompts[agentName] = make(map[string]PromptVersion)
	}

	if _, exists := pm.prompts[agentName][version.Version]; exists {
		return fmt.Errorf("prompt version %s for agent %s already exists", version.Version, agentName)
	}

	pm.prompts[agentName][version.Version] = version
	return nil
}

// GetPrompt retrieves a prompt for a specific agent and version
func (pm *PromptManager) GetPrompt(agentName, version string) (PromptVersion, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	versions, exists := pm.prompts[agentName]
	if !exists {
		return PromptVersion{}, fmt.Errorf("no prompts found for agent %s", agentName)
	}

	prompt, exists := versions[version]
	if !exists {
		return PromptVersion{}, fmt.Errorf("prompt version %s not found for agent %s", version, agentName)
	}

	return prompt, nil
}

// GetLatestVersion retrieves the latest prompt version for an agent
func (pm *PromptManager) GetLatestVersion(agentName string) (PromptVersion, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	versions, exists := pm.prompts[agentName]
	if !exists {
		return PromptVersion{}, fmt.Errorf("no prompts found for agent %s", agentName)
	}

	var latest PromptVersion
	for _, version := range versions {
		if latest.Version == "" || version.Version > latest.Version {
			latest = version
		}
	}

	return latest, nil
}

// ListVersions returns all prompt versions for an agent
func (pm *PromptManager) ListVersions(agentName string) ([]PromptVersion, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	versions, exists := pm.prompts[agentName]
	if !exists {
		return nil, fmt.Errorf("no prompts found for agent %s", agentName)
	}

	result := make([]PromptVersion, 0, len(versions))
	for _, version := range versions {
		result = append(result, version)
	}

	return result, nil
}

// Prompt represents a prompt
type Prompt struct {
	system  string
	user    string
	version string
}

// NewPrompt creates a new Prompt instance
func NewPrompt(system, user, version string) core.Prompt {
	return &Prompt{
		system:  system,
		user:    user,
		version: version,
	}
}

// GetSystemPrompt returns the system prompt
func (p *Prompt) GetSystemPrompt() string {
	return p.system
}

// GetUserPrompt returns the user prompt
func (p *Prompt) GetUserPrompt() string {
	return p.user
}

// GetSystemPromptByRound returns the system prompt for a specific round
func (p *Prompt) GetSystemPromptByRound(round int) string {
	return p.system
}

// GetUserPromptByRound returns the user prompt for a specific round
func (p *Prompt) GetUserPromptByRound(round int) string {
	return p.user
}

// GetVersion returns the prompt version
func (p *Prompt) GetVersion() string {
	return p.version
}
