package cache

import (
	"strings"
)

type Key string

const (
	SmsCodeKey Key = "user:sms:code:phone"
)

func (k Key) Key(params ...string) string {
	var key strings.Builder
	key.WriteString(string(k))
	for _, param := range params {
		key.WriteString(":")
		key.WriteString(param)
	}
	return key.String()
}
