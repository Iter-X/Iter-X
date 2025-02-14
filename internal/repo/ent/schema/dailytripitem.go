package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DailyTripItem holds the schema definition for the DailyTripItem entity.
type DailyTripItem struct {
	ent.Schema
}

// Fields of the DailyTripItem.
func (DailyTripItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.UUID("trip_id", uuid.UUID{}).Unique(),
		field.UUID("daily_trip_id", uuid.UUID{}).Unique(),
		field.String("notes").MaxLen(255).Optional(),
	}
}

// Edges of the DailyTripItem.
func (DailyTripItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("trip", Trip.Type).
			Ref("daily_trip_item").
			Field("trip_id").
			Unique().Required(),
		edge.From("daily_trip", DailyTrip.Type).
			Ref("daily_trip_item").
			Field("daily_trip_id").
			Unique().Required(),
	}
}

func (DailyTripItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
