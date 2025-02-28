package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DailyTrip holds the schema definition for the DailyTrip entity.
type DailyTrip struct {
	ent.Schema
}

// Fields of the DailyTrip.
func (DailyTrip) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.UUID("trip_id", uuid.UUID{}).Unique(),
		field.Int32("day").Positive(),
		field.Time("date"),
		field.String("notes").MaxLen(255).Optional(),
	}
}

// Edges of the DailyTrip.
func (DailyTrip) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("trip", Trip.Type).
			Ref("daily_trip").
			Field("trip_id").
			Unique().Required(),
		edge.To("daily_itinerary", DailyItinerary.Type),
	}
}

func (DailyTrip) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
