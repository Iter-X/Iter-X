package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// TimeMixin implements the ent.Mixin for sharing
// time fields with package schemas.
type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

type OperatorMixin struct {
	mixin.Schema
}

func (OperatorMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("updated_by", uuid.UUID{}),
		field.UUID("created_by", uuid.UUID{}),
	}
}

type LocalizedNameMixin struct {
	mixin.Schema
}

func (LocalizedNameMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name_local").NotEmpty().MaxLen(255),
		field.String("name_en").MaxLen(255),
		field.String("name_cn").MaxLen(255),
	}
}
