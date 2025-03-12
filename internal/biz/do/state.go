package do

import (
	"time"
)

// State is the model entity for the State schema.
type State struct {
	// ID of the ent.
	ID uint `json:"id,omitempty"`
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
	// state code
	Code string `json:"code,omitempty"`
	// country id
	CountryID uint `json:"country_id,omitempty"`

	// Poi holds the value of the poi edge.
	Poi []*PointsOfInterest `json:"poi,omitempty"`
	// City holds the value of the city edge.
	City []*City `json:"city,omitempty"`
	// Country holds the value of the country edge.
	Country *Country `json:"country,omitempty"`
}
