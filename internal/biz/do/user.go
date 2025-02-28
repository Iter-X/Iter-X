package do

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// AvatarURL holds the value of the "avatar_url" field.
	AvatarURL string `json:"avatar_url,omitempty"`

	// RefreshToken holds the value of the refresh_token edge.
	RefreshTokens []*RefreshToken `json:"refresh_token,omitempty"`
	// Trip holds the value of the trip edge.
	Trips []*Trip `json:"trip,omitempty"`
}
