package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Language holds the schema definition for the Language entity.
type Language struct {
	ent.Schema
}

// Fields of the Language.
func (Language) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Unique().MaxLen(10),
		field.String("name").MaxLen(50),
		field.String("native_name").MaxLen(50),
		field.Bool("enabled").Default(true),
	}
}

func (Language) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
