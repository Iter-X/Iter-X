package sms_test

import (
	"context"
	"testing"

	"github.com/iter-x/iter-x/pkg/sms"
	"github.com/iter-x/iter-x/pkg/xcontext"
)

var _ sms.ClientConfig = (*Config)(nil)

type Config struct {
	accessKeyId     string
	accessKeySecret string
}

func (c *Config) GetAccessKeyId() string {
	return c.accessKeyId
}

func (c *Config) GetAccessKeySecret() string {
	return c.accessKeySecret
}

type AuthTokensConfig struct {
	bundleId        string
	expire          int64
	packageName     string
	signName        string
	sceneCode       string
	smsTemplateCode string
	smsCodeExpire   int32
}

func (a *AuthTokensConfig) GetBundleId() string {
	return a.bundleId
}

func (a *AuthTokensConfig) GetExpire() int64 {
	return a.expire
}

func (a *AuthTokensConfig) GetPackageName() string {
	return a.packageName
}

func (a *AuthTokensConfig) GetSignName() string {
	return a.signName
}

func (a *AuthTokensConfig) GetSceneCode() string {
	return a.sceneCode
}

func (a *AuthTokensConfig) GetSmsTemplateCode() string {
	return a.smsTemplateCode
}

func (a *AuthTokensConfig) GetSmsCodeExpire() int32 {
	return a.smsCodeExpire
}

var _ sms.AuthTokensConfig = (*AuthTokensConfig)(nil)

func TestNewClient(t *testing.T) {
	conf := &Config{
		accessKeyId:     "accessKeyId",
		accessKeySecret: "accessKeySecret",
	}
	c := sms.NewClient(sms.WithClientConfig(conf))
	ctx := xcontext.WithClientType(context.Background(), "Android")

	authTokensConfig := &AuthTokensConfig{
		bundleId:        "bundleId",
		expire:          100,
		packageName:     "packageName",
		signName:        "signName",
		sceneCode:       "sceneCode",
		smsTemplateCode: "smsTemplateCode",
		smsCodeExpire:   100,
	}
	tokens, err := c.GetSmsAuthTokens(ctx, authTokensConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tokens.GetData())
}
