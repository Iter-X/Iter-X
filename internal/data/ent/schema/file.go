package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique().Comment("file id"),
		field.UUID("user_id", uuid.UUID{}).Comment("user id"),
		field.String("name").MaxLen(255).Comment("file name"),
		field.String("object_key").MaxLen(255).Comment("file object key"),
		field.Uint("size").Optional().Comment("file size in bytes"),
		field.String("ext").MaxLen(255).Comment("file extension"),
	}
}

func (File) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "files"},
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("files").Unique().Required().Field("user_id"),
		edge.To("poi_files", PointsOfInterestFiles.Type),
	}
}

func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (File) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("object_key").Unique(),
	}
}
