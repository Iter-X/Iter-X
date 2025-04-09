package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Agent holds the schema definition for the Agent entity.
type Agent struct {
	ent.Schema
}

// Fields of the Agent.
func (Agent) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique().
			Comment("Agent name"),
		field.Bool("enabled").
			Default(true).
			Comment("Whether the agent is enabled"),
		field.String("description").
			Optional().
			Comment("Agent description"),
	}
}

// Edges of the Agent.
func (Agent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("prompts", AgentPrompt.Type).
			Comment("Agent's prompt versions"),
		edge.To("tool_bindings", AgentToolBinding.Type).
			Comment("Tool bindings for this agent"),
	}
}

// Mixin of the Agent.
func (Agent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
