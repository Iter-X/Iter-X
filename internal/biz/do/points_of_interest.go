package do

import (
	"time"

	"github.com/google/uuid"
)

// PointsOfInterest is the model entity for the PointsOfInterest schema.
type PointsOfInterest struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// NameEn holds the value of the "name_en" field.
	NameEn string `json:"name_en,omitempty"`
	// NameCn holds the value of the "name_cn" field.
	NameCn    string `json:"name_cn,omitempty"`
	NameLocal string `json:"name_local,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Latitude holds the value of the "latitude" field.
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude holds the value of the "longitude" field.
	Longitude float64 `json:"longitude,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Category holds the value of the "category" field.
	Category string `json:"category,omitempty"`
	// Rating holds the value of the "rating" field.
	Rating float32 `json:"rating,omitempty"`
	// RecommendedDurationMinutes holds the value of the "recommended_duration_minutes" field.
	RecommendedDurationMinutes int64 `json:"recommended_duration_minutes,omitempty"`
	// CityID holds the value of the "city_id" field.
	CityID uint `json:"city_id,omitempty"`
	// StateID holds the value of the "state_id" field.
	StateID uint `json:"state_id,omitempty"`
	// CountryID holds the value of the "country_id" field.
	CountryID uint `json:"country_id,omitempty"`
	// ContinentID holds the value of the "continent_id" field.
	ContinentID uint `json:"continent_id,omitempty"`

	// City holds the value of the city edge.
	City *City `json:"city,omitempty"`
	// State holds the value of the state edge.
	State *State `json:"state,omitempty"`
	// Country holds the value of the country edge.
	Country *Country `json:"country,omitempty"`
	// Continent holds the value of the continent edge.
	Continent *Continent `json:"continent,omitempty"`
	// DailyItinerary holds the value of the daily_itinerary edge.
	DailyItinerary []*DailyItinerary `json:"daily_itinerary,omitempty"`
	// PoiFiles holds the value of the poi_files edge.
	PoiFiles []*PointsOfInterestFiles `json:"poi_files,omitempty"`
}
