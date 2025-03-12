package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Continent holds the schema definition for the Continent entity.
type Continent struct {
	ent.Schema
}

// Fields of the Continent.
func (Continent) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),
		field.String("code").MaxLen(20).Comment("continent code"),
	}
}

// Edges of the Continent.
func (Continent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("poi", PointsOfInterest.Type),
		edge.To("country", Country.Type),
	}
}

func (Continent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		LocalizedNameMixin{},
	}
}
