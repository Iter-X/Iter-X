package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// State holds the schema definition for the State entity.
type State struct {
	ent.Schema
}

// Fields of the State.
func (State) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
	}
}

// Edges of the State.
func (State) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("poi", PointsOfInterest.Type),
	}
}

func (State) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		LocalizedNameMixin{},
	}
}
