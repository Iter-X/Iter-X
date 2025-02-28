package do

import (
	"time"

	"github.com/google/uuid"
)

// City is the model entity for the City schema.
type City struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// NameEn holds the value of the "name_en" field.
	NameEn string `json:"name_en,omitempty"`
	// NameCn holds the value of the "name_cn" field.
	NameCn string `json:"name_cn,omitempty"`

	// Poi holds the value of the poi edge.
	Poi []*PointsOfInterest `json:"poi,omitempty"`
}
