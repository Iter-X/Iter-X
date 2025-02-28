package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Country holds the schema definition for the Country entity.
type Country struct {
	ent.Schema
}

// Fields of the Country.
func (Country) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
	}
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("poi", PointsOfInterest.Type),
	}
}

func (Country) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		LocalizedNameMixin{},
	}
}
