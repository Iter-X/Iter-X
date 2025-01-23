package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RefreshToken holds the schema definition for the RefreshToken entity.
type RefreshToken struct {
	ent.Schema
}

// Fields of the RefreshToken.
func (RefreshToken) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.String("token").NotEmpty().MaxLen(64),
		field.Time("expires_at"),
		field.UUID("user_id", uuid.UUID{}).Unique(),
	}
}

// Edges of the RefreshToken.
func (RefreshToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("refresh_token").
			Field("user_id").
			Unique().Required(),
	}
}

func (RefreshToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
