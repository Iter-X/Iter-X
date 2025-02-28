package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// City holds the schema definition for the City entity.
type City struct {
	ent.Schema
}

// Fields of the City.
func (City) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
	}
}

// Edges of the City.
func (City) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("poi", PointsOfInterest.Type),
	}
}

func (City) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		LocalizedNameMixin{},
	}
}
