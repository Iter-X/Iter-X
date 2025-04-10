package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TripPOIPool holds the schema definition for the TripPOIPool entity.
type TripPOIPool struct {
	ent.Schema
}

func (TripPOIPool) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "trip_poi_pool"},
	}
}

// Fields of the TripPOIPool.
func (TripPOIPool) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.UUID("trip_id", uuid.UUID{}),
		field.UUID("poi_id", uuid.UUID{}),
		field.String("notes").Optional(),
	}
}

// Edges of the TripPOIPool.
func (TripPOIPool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("trip", Trip.Type).Ref("poi_pool").Field("trip_id").Unique().Required(),
		edge.From("poi", PointsOfInterest.Type).Ref("trip_pools").Field("poi_id").Unique().Required(),
	}
}

func (TripPOIPool) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		OperatorMixin{},
	}
}
