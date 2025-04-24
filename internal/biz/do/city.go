package do

import (
	"time"
)

// City is the model entity for the City schema.
type City struct {
	// city id
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
	// city code
	Code string `json:"code,omitempty"`
	// state id
	StateID uint `json:"state_id,omitempty"`

	// Poi holds the value of the poi edge.
	Poi []*PointsOfInterest `json:"poi,omitempty"`
	// State holds the value of the state edge.
	State *State `json:"state,omitempty"`
}

func (c *City) GetNameCn() string {
	if c == nil {
		return ""
	}
	return c.NameCn
}
