package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PointsOfInterest holds the schema definition for the PointsOfInterest entity.
type PointsOfInterest struct {
	ent.Schema
}

func (PointsOfInterest) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "points_of_interest"},
	}
}

// Fields of the PointsOfInterest.
func (PointsOfInterest) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.String("description").MaxLen(1000),
		field.String("address").MaxLen(255),
		field.Float("latitude"),
		field.Float("longitude"),
		field.String("type").MaxLen(50),     // Attractions, Restaurants, Hotels, Cities, etc.
		field.String("category").MaxLen(50), // Historical, Modern, Natural, etc.
		field.Float32("rating").Positive(),
		field.Int64("recommended_duration_minutes").Positive(),
		field.UUID("city_id", uuid.UUID{}).Optional(),
		field.UUID("state_id", uuid.UUID{}).Optional(),
		field.UUID("country_id", uuid.UUID{}).Optional(),
		field.UUID("continent_id", uuid.UUID{}).Optional(),
	}
}

// Edges of the PointsOfInterest.
func (PointsOfInterest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("city", City.Type).Ref("poi").Field("city_id").Unique(),
		edge.From("state", State.Type).Ref("poi").Field("state_id").Unique(),
		edge.From("country", Country.Type).Ref("poi").Field("country_id").Unique(),
		edge.From("continent", Continent.Type).Ref("poi").Field("continent_id").Unique(),
		edge.To("daily_itinerary", DailyItinerary.Type),
	}
}

func (PointsOfInterest) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		LocalizedNameMixin{},
	}
}
