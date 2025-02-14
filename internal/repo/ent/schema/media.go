package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/common/model"
)

// Media holds the schema definition for the Media entity.
type Media struct {
	ent.Schema
}

// Fields of the Media.
func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.String("filename").NotEmpty().MaxLen(255),
		field.String("file_type").NotEmpty().MaxLen(255),
		field.Uint8("storage_type").GoType(model.StorageType(0)),
		field.String("path").NotEmpty().MaxLen(255),
	}
}

// Edges of the Media.
func (Media) Edges() []ent.Edge {
	return nil
}

func (Media) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
