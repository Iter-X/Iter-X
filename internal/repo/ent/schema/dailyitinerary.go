package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DailyItinerary holds the schema definition for the DailyItinerary entity.
type DailyItinerary struct {
	ent.Schema
}

// Fields of the DailyItinerary.
func (DailyItinerary) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.UUID("trip_id", uuid.UUID{}).Unique(),
		field.UUID("daily_trip_id", uuid.UUID{}).Unique(),
		field.UUID("poi_id", uuid.UUID{}),
		field.String("notes").MaxLen(255).Optional(),
	}
}

// Edges of the DailyItinerary.
func (DailyItinerary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("trip", Trip.Type).
			Ref("daily_itinerary").
			Field("trip_id").
			Unique().Required(),
		edge.From("daily_trip", DailyTrip.Type).
			Ref("daily_itinerary").
			Field("daily_trip_id").
			Unique().Required(),
		edge.From("poi", PointsOfInterest.Type).
			Ref("daily_itinerary").
			Field("poi_id").
			Unique().Required(),
	}
}

func (DailyItinerary) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
