package vobj

//go:generate go run ../../cmd/stringer/cmd.go -type=ClientType -linecomment -output=client_type.string.go
type ClientType int

const (
	ClientTypeUnknown ClientType = iota // unknown
	ClientTypeAndroid                   // Android
	ClientTypeIOS                       // iOS
)

// BuildClientType builds a ClientType from a string.
func BuildClientType(s string) ClientType {
	switch s {
	case "Android", "android":
		return ClientTypeAndroid
	case "iOS", "ios":
		return ClientTypeIOS
	default:
		return ClientTypeUnknown
	}
}
