package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DailyTripLocation holds the schema definition for the DailyTripLocation entity.
type DailyTripLocation struct {
	ent.Schema
}

// Fields of the DailyTripLocation.
func (DailyTripLocation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("daily_trip_id", uuid.UUID{}),
		field.UUID("location_id", uuid.UUID{}),
		field.String("location_type").
			Comment("Type of location: Continent, Country, State, City, POI"),
		field.Int8("sequence").
			Comment("Order of locations within the daily trip"),
	}
}

// Edges of the DailyTripLocation.
func (DailyTripLocation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("daily_trip", DailyTrip.Type).
			Ref("daily_trip_locations").
			Field("daily_trip_id").
			Unique().
			Required(),
	}
}

func (DailyTripLocation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
