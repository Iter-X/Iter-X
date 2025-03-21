package vobj

import (
	"strings"
)

//go:generate stringer -type=ClientType -linecomment -output=client_type.string.go
type ClientType int

const (
	ClientTypeUnknown ClientType = iota // unknown
	ClientTypeAndroid                   // Android
	ClientTypeIOS                       // iOS
)

// BuildClientType builds a ClientType from a string.
func BuildClientType(s string) ClientType {
	switch strings.ToLower(s) {
	case "android":
		return ClientTypeAndroid
	case "ios":
		return ClientTypeIOS
	default:
		return ClientTypeUnknown
	}
}
