package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Trip holds the schema definition for the Trip entity.
type Trip struct {
	ent.Schema
}

// Fields of the Trip.
func (Trip) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}).Unique(),
		field.Bool("status").Default(true),
		field.String("title").NotEmpty().MaxLen(50),
		field.String("description").MaxLen(255),
		field.Time("start_date"),
		field.Time("end_date"),
		field.Int("days").NonNegative(),
	}
}

// Edges of the Trip.
func (Trip) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("trip").
			Field("user_id").
			Unique().Required(),
		edge.To("daily_trip", DailyTrip.Type),
		edge.To("daily_itinerary", DailyItinerary.Type),
		edge.To("collaborators", User.Type),
	}
}

func (Trip) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
