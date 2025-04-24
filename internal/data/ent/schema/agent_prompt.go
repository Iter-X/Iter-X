package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AgentPrompt holds the schema definition for the AgentPrompt entity.
type AgentPrompt struct {
	ent.Schema
}

// Fields of the AgentPrompt.
func (AgentPrompt) Fields() []ent.Field {
	return []ent.Field{
		field.String("version").
			NotEmpty().
			Comment("Prompt version"),
		field.String("system").
			NotEmpty().
			Comment("System prompt message"),
		field.String("user").
			NotEmpty().
			Comment("User prompt message"),
	}
}

// Edges of the AgentPrompt.
func (AgentPrompt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("agent", Agent.Type).
			Ref("prompts").
			Unique().
			Required().
			Comment("The agent this prompt belongs to"),
	}
}

// Mixin of the AgentPrompt.
func (AgentPrompt) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
