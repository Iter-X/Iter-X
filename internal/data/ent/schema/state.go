package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// State holds the schema definition for the State entity.
type State struct {
	ent.Schema
}

// Fields of the State.
func (State) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),
		field.String("code").MaxLen(20).Comment("state code"),
		field.Uint("country_id").Comment("country id"),
	}
}

// Edges of the State.
func (State) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("poi", PointsOfInterest.Type),
		edge.To("city", City.Type),
		edge.From("country", Country.Type).Ref("state").Unique().Required().Field("country_id"),
	}
}

func (State) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		LocalizedNameMixin{},
	}
}
