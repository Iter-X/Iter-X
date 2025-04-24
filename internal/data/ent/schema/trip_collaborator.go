package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/data/cnst"
)

// TripCollaborator holds the schema definition for the TripCollaborator entity.
type TripCollaborator struct {
	ent.Schema
}

// Fields of the TripCollaborator.
func (TripCollaborator) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.UUID("trip_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.Enum("status").
			Values(cnst.CollaboratorStatusInvited, cnst.CollaboratorStatusAccepted, cnst.CollaboratorStatusRejected).
			Default(cnst.CollaboratorStatusInvited),
	}
}

// Edges of the TripCollaborator.
func (TripCollaborator) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("trip", Trip.Type).Ref("trip_collaborators").Field("trip_id").Unique().Required(),
		edge.From("user", User.Type).Ref("trip_collaborators").Field("user_id").Unique().Required(),
	}
}

func (TripCollaborator) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
