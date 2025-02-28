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

func TestNewClient(t *testing.T) {
	conf := &Config{
		accessKeyId:     "accessKeyId",
		accessKeySecret: "accessKeySecret",
	}
	c := sms.NewClient(sms.WithClientConfig(conf))
	ctx := xcontext.WithClientType(context.Background(), "Android")

	c.GetSmsAuthTokens()

}
