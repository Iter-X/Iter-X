package do

import (
	"time"

	"github.com/google/uuid"
)

type TripCollaborator struct {
	ID        uuid.UUID
	TripID    uuid.UUID
	UserID    uuid.UUID
	Status    string
	Username  string
	Nickname  string
	AvatarURL string
	CreatedAt time.Time
	UpdatedAt time.Time
}
