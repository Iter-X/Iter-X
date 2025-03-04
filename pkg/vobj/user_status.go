package vobj

//go:generate go run ../../cmd/stringer/cmd.go -type=UserStatus -linecomment -output=user_status.string.go
type UserStatus int

const (
	UserStatusUnknown  UserStatus = iota // Unknown
	UserStatusActive                     // Active
	UserStatusInactive                   // Inactive
)
