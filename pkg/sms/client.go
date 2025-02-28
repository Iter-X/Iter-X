package sms

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dypnsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"go.uber.org/zap"
)

type (
	Client struct {
		accessKeyId     string
		accessKeySecret string

		client *dysmsapi.Client

		logger *zap.SugaredLogger
	}

	ClientConfig interface {
		GetAccessKeyId() string
		GetAccessKeySecret() string
	}

	ClientOption func(c *Client)
)

// NewClient creates a new SMS client
func NewClient(opts ...ClientOption) *Client {
	c := &Client{}
	for _, opt := range opts {
		opt(c)
	}
	c.client = c.initClient()
	return c
}

// WithClientConfig sets the client config
func WithClientConfig(cfg ClientConfig) ClientOption {
	return func(c *Client) {
		c.accessKeyId = cfg.GetAccessKeyId()
		c.accessKeySecret = cfg.GetAccessKeySecret()
	}
}

// WithLogger sets the logger
func WithLogger(logger *zap.SugaredLogger) ClientOption {
	return func(c *Client) {
		c.logger = logger
	}
}

// initClient initializes the SMS client
func (c *Client) initClient() *dysmsapi.Client {
	if c.accessKeySecret == "" || c.accessKeyId == "" {
		panic("SMS sending credential information is not configured")
	}
	config := &openapi.Config{
		AccessKeyId:     &c.accessKeyId,
		AccessKeySecret: &c.accessKeySecret,
	}
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	client, err := dysmsapi.NewClient(config)
	if err != nil {
		panic(err)
	}
	return client
}
