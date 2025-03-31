package do

import (
	"time"
)

// Continent is the model entity for the Continent schema.
type Continent struct {
	// ID of the ent.
	ID uint `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// NameEn holds the value of the "name_en" field.
	NameEn string `json:"name_en,omitempty"`
	// NameCn holds the value of the "name_cn" field.
	NameCn    string `json:"name_cn,omitempty"`
	NameLocal string `json:"name_local,omitempty"`
	// continent code
	Code string `json:"code,omitempty"`

	// Poi holds the value of the poi edge.
	Poi []*PointsOfInterest `json:"poi,omitempty"`
	// Country holds the value of the country edge.
	Country []*Country `json:"country,omitempty"`
}
