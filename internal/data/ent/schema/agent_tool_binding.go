package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AgentToolBinding holds the schema definition for the AgentToolBinding entity.
type AgentToolBinding struct {
	ent.Schema
}

// Fields of the AgentToolBinding.
func (AgentToolBinding) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("enabled").
			Default(true).
			Comment("Whether this tool is enabled for the agent"),
	}
}

// Edges of the AgentToolBinding.
func (AgentToolBinding) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("agent", Agent.Type).
			Ref("tool_bindings").
			Unique().
			Required().
			Comment("The agent in this binding"),
		edge.From("tool", AgentTool.Type).
			Ref("agent_bindings").
			Unique().
			Required().
			Comment("The tool in this binding"),
	}
}

// Mixin of the AgentToolBinding.
func (AgentToolBinding) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
