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

type PointsOfInterestFiles struct {
	ent.Schema
}

func (PointsOfInterestFiles) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique().Comment("id"),
		field.UUID("poi_id", uuid.UUID{}).Comment("poi id"),
		field.Uint("file_id").Comment("file id"),
	}
}

func (PointsOfInterestFiles) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "poi_files"},
	}
}

func (PointsOfInterestFiles) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("poi", PointsOfInterest.Type).Ref("poi_files").Unique().Required().Field("poi_id"),
		edge.From("file", File.Type).Ref("poi_files").Unique().Required().Field("file_id"),
	}
}

func (PointsOfInterestFiles) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (PointsOfInterestFiles) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("poi_id", "file_id").Unique(),
	}
}
