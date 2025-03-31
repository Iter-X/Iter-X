package do

import (
	"time"
)

// Country is the model entity for the Country schema.
type Country struct {
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
	// country code
	Code string `json:"code,omitempty"`
	// continent id
	ContinentID uint `json:"continent_id,omitempty"`

	// Poi holds the value of the poi edge.
	Poi []*PointsOfInterest `json:"poi,omitempty"`
	// State holds the value of the state edge.
	State []*State `json:"state,omitempty"`
	// Continent holds the value of the continent edge.
	Continent *Continent `json:"continent,omitempty"`
}
