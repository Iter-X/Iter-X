package conf

func (x *Auth_AuthTokens) GetSmsCodeExpireSec() int32 {
	t := x.GetSmsCodeExpire()
	if t == nil {
		return 0
	}
	return int32(t.GetSeconds())
}

func (x *Auth_AuthTokens) GetExpireSec() int64 {
	t := x.GetExpire()
	if t == nil {
		return 0
	}
	return t.GetSeconds()
}
