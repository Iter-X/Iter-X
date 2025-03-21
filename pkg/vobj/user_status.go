package vobj

//go:generate stringer -type=UserStatus -linecomment -output=user_status.string.go
type UserStatus int8

const (
	UserStatusUnknown  UserStatus = iota // Unknown
	UserStatusActive                     // Active
	UserStatusInactive                   // Inactive
)
