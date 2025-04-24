package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserPreference holds the schema definition for the UserPreference entity.
type UserPreference struct {
	ent.Schema
}

// Fields of the UserPreference.
func (UserPreference) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.String("app_language").Default("en").MaxLen(10),
		field.String("default_city").Optional().MaxLen(50),
		field.Enum("time_format").Values("12h", "24h").Default("24h"),
		field.Enum("distance_unit").Values("km", "mile").Default("km"),
		field.Enum("dark_mode").Values("on", "off", "system").Default("system"),
		field.Bool("notify_itinerary").Default(true),
		field.Bool("notify_community").Default(true),
		field.Bool("notify_recommendations").Default(true),
	}
}

// Edges of the UserPreference.
func (UserPreference) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("preference").
			Unique().
			Required().
			Field("user_id"),
	}
}

func (UserPreference) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
