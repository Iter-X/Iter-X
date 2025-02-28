package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Continent holds the schema definition for the Continent entity.
type Continent struct {
	ent.Schema
}

// Fields of the Continent.
func (Continent) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
	}
}

// Edges of the Continent.
func (Continent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("poi", PointsOfInterest.Type),
	}
}

func (Continent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		LocalizedNameMixin{},
	}
}
