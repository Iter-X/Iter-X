package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/iter-x/iter-x/pkg/vobj"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.Int8("status").Default(vobj.UserStatusActive.GetValue()),
		field.String("username").NotEmpty().MaxLen(50),
		field.String("password").NotEmpty().MaxLen(255),
		field.String("salt").NotEmpty().MaxLen(255),
		field.String("nickname").Optional().MaxLen(50),
		field.String("remark").Optional().MaxLen(50),
		field.String("phone").MaxLen(11),
		field.String("email").MaxLen(255),
		field.String("avatar_url").MaxLen(255).Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("refresh_token", RefreshToken.Type),
		edge.To("trip", Trip.Type),
		edge.From("collaborated_trips", Trip.Type).Ref("collaborators"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
