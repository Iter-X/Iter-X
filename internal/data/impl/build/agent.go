package build

import (
	"encoding/json"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

// ConvertFunctionSpec converts a map[string]interface{} to *do.FunctionConfig
func ConvertFunctionSpec(spec map[string]interface{}) *do.FunctionConfig {
	if spec == nil {
		return nil
	}

	// Convert the map to JSON
	jsonBytes, err := json.Marshal(spec)
	if err != nil {
		return nil
	}

	// Convert JSON to FunctionConfig
	var config do.FunctionConfig
	if err := json.Unmarshal(jsonBytes, &config); err != nil {
		return nil
	}

	return &config
}

// ConvertFunctionConfig converts a *do.FunctionConfig to map[string]interface{}
func ConvertFunctionConfig(config *do.FunctionConfig) map[string]interface{} {
	if config == nil {
		return nil
	}

	// Convert the config to JSON
	jsonBytes, err := json.Marshal(config)
	if err != nil {
		return nil
	}

	// Convert JSON to map
	var spec map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &spec); err != nil {
		return nil
	}

	return spec
}

// AgentRepositoryImplToEntity converts an agent entity to a business object
func AgentRepositoryImplToEntity(po *ent.Agent) *do.Agent {
	if po == nil {
		return nil
	}

	agent := &do.Agent{
		ID:          po.ID,
		CreatedAt:   po.CreatedAt,
		UpdatedAt:   po.UpdatedAt,
		Name:        po.Name,
		Enabled:     po.Enabled,
		Description: po.Description,
	}

	// Convert prompt
	if len(po.Edges.Prompts) > 0 {
		prompt := po.Edges.Prompts[0]
		agent.Prompt = &do.Prompt{
			ID:        prompt.ID,
			CreatedAt: prompt.CreatedAt,
			UpdatedAt: prompt.UpdatedAt,
			Version:   prompt.Version,
			System:    prompt.System,
			User:      prompt.User,
		}
	}

	// Convert tools
	if len(po.Edges.ToolBindings) > 0 {
		agent.Tools = make([]*do.Tool, 0, len(po.Edges.ToolBindings))
		for _, binding := range po.Edges.ToolBindings {
			tool := binding.Edges.Tool
			agent.Tools = append(agent.Tools, &do.Tool{
				ID:          tool.ID,
				CreatedAt:   tool.CreatedAt,
				UpdatedAt:   tool.UpdatedAt,
				Name:        tool.Name,
				Type:        tool.Type,
				Enabled:     binding.Enabled,
				Description: tool.Description,
				BaseURL:     tool.BaseURL,
				APIKey:      tool.APIKey,
				Model:       tool.Model,
				Function:    ConvertFunctionSpec(tool.FunctionSpec),
			})
		}
	}

	return agent
}

// AgentRepositoryImplToEntities converts agent entities to business objects
func AgentRepositoryImplToEntities(pos []*ent.Agent) []*do.Agent {
	if pos == nil {
		return nil
	}

	agents := make([]*do.Agent, 0, len(pos))
	for _, po := range pos {
		agents = append(agents, AgentRepositoryImplToEntity(po))
	}
	return agents
}
