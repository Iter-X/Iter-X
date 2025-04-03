package do

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	// file id
	ID uint `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// user id
	UserID uuid.UUID `json:"user_id,omitempty"`
	// file name
	Name string `json:"name,omitempty"`
	// file object key
	ObjectKey string `json:"object_key,omitempty"`
	// file size in bytes
	Size uint `json:"size,omitempty"`
	// file extension
	Ext string `json:"ext,omitempty"`

	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// PoiFiles holds the value of the poi_files edge.
	PoiFiles []*PointsOfInterestFiles `json:"poi_files,omitempty"`
}

type PointsOfInterestFiles struct {
	// ID of the ent.
	// id
	ID uint `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// poi id
	PoiID uuid.UUID `json:"poi_id,omitempty"`
	// file id
	FileID uint `json:"file_id,omitempty"`

	// Poi holds the value of the poi edge.
	Poi *PointsOfInterest `json:"poi,omitempty"`
	// File holds the value of the file edge.
	File *File `json:"file,omitempty"`
}
