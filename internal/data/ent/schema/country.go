package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Country holds the schema definition for the Country entity.
type Country struct {
	ent.Schema
}

// Fields of the Country.
func (Country) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),
		field.String("code").MaxLen(20).Comment("country code"),
		field.Uint("continent_id").Comment("continent id"),
		field.Uint("file_id").Optional().Comment("country image file id"),
	}
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("poi", PointsOfInterest.Type),
		edge.To("state", State.Type),
		edge.From("continent", Continent.Type).Ref("country").Unique().Required().Field("continent_id"),
		edge.To("daily_trip_location", DailyTripLocation.Type),
		edge.To("image", File.Type).Unique().Field("file_id"),
	}
}

func (Country) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		LocalizedNameMixin{},
	}
}
