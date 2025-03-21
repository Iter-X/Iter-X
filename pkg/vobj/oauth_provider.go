package vobj

//go:generate stringer -type=OAuthProvider -linecomment -output=oauth_provider.string.go
type OAuthProvider int

const (
	OAuthProviderGOOGLE OAuthProvider = iota // GOOGLE
	OAuthProviderGITHUB                      // GITHUB
	OAuthProviderWECHAT                      // WECHAT
)
