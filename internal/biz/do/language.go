package do

import "time"

type Language struct {
	Code       string
	Name       string
	NativeName string
	Enabled    bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
