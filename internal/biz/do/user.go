package do

import (
	"time"

	"github.com/google/uuid"
	"github.com/iter-x/iter-x/pkg/vobj"
)

type User struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status holds the value of the "status" field.
	Status vobj.UserStatus `json:"status,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Salt holds the value of the "salt" field.
	Salt string `json:"salt,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// AvatarURL holds the value of the "avatar_url" field.
	AvatarURL string `json:"avatar_url,omitempty"`

	// RefreshToken holds the value of the refresh_token edge.
	RefreshTokens []*RefreshToken `json:"refresh_token,omitempty"`
	// Trip holds the value of the trip edge.
	Trips []*Trip `json:"trip,omitempty"`
}

type UserPreference struct {
	ID                    uuid.UUID
	UserID                uuid.UUID
	AppLanguage           string
	DefaultCity           string
	TimeFormat            string
	DistanceUnit          string
	DarkMode              string
	TripReminder          bool
	CommunityNotification bool
	RecommendContentPush  bool
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
