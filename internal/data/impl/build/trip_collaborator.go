package build

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func TripCollaboratorRepositoryImplToEntity(po *ent.TripCollaborator) *do.TripCollaborator {
	if po == nil {
		return nil
	}

	collaborator := &do.TripCollaborator{
		ID:        po.ID,
		TripID:    po.TripID,
		UserID:    po.UserID,
		Status:    po.Status.String(),
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
	}

	// Add user information if available
	if po.Edges.User != nil {
		collaborator.Username = po.Edges.User.Username
		collaborator.Nickname = po.Edges.User.Nickname
		collaborator.AvatarURL = po.Edges.User.AvatarURL
	}

	return collaborator
}

func TripCollaboratorRepositoryImplToEntities(pos []*ent.TripCollaborator) []*do.TripCollaborator {
	if pos == nil {
		return nil
	}

	entities := make([]*do.TripCollaborator, 0, len(pos))
	for _, po := range pos {
		if entity := TripCollaboratorRepositoryImplToEntity(po); entity != nil {
			entities = append(entities, entity)
		}
	}
	return entities
}
