package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AgentTool holds the schema definition for the AgentTool entity.
type AgentTool struct {
	ent.Schema
}

// Fields of the AgentTool.
func (AgentTool) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique().
			Comment("Tool name"),
		field.String("description").
			Optional().
			Comment("Tool description"),
		field.String("type").
			NotEmpty().
			Comment("Tool type (e.g. CodeUse, LLMUse)"),
		field.String("base_url").
			Optional().
			Comment("Base URL for API tools"),
		field.String("api_key").
			Optional().
			Sensitive().
			Comment("API key for external services"),
		field.String("model").
			Optional().
			Comment("Model name for LLM tools"),
		field.JSON("function_spec", map[string]interface{}{}).
			Optional().
			Comment("Function specification for LLMUse tools"),
	}
}

// Edges of the AgentTool.
func (AgentTool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("agent_bindings", AgentToolBinding.Type).
			Comment("Agent bindings for this tool"),
	}
}

// Mixin of the AgentTool.
func (AgentTool) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
