package vobj

//go:generate go run ../../cmd/stringer/cmd.go -type=OAuthProvider -linecomment -output=oauth_provider.string.go
type OAuthProvider int

const (
	OAuthProviderGOOGLE OAuthProvider = iota // GOOGLE
	OAuthProviderGITHUB                      // GITHUB
	OAuthProviderWECHAT                      // WECHAT
)
