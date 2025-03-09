package sms

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapiV2 "github.com/alibabacloud-go/dypnsapi-20170525/v2/client"
	dysmsapiV3 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"go.uber.org/zap"
)

type (
	Client struct {
		accessKeyId     string
		accessKeySecret string
		endpoint        string

		clientV2 *dysmsapiV2.Client
		clientV3 *dysmsapiV3.Client

		logger *zap.SugaredLogger
	}

	ClientConfig interface {
		GetAccessKeyId() string
		GetAccessKeySecret() string
		GetEndpoint() string
	}

	ClientOption func(c *Client)
)

// NewClient creates a new SMS client
func NewClient(opts ...ClientOption) *Client {
	c := &Client{}
	for _, opt := range opts {
		opt(c)
	}
	c.clientV2 = c.initClientV2()
	c.clientV3 = c.initClientV3()
	return c
}

// WithClientConfig sets the client config
func WithClientConfig(cfg ClientConfig) ClientOption {
	return func(c *Client) {
		c.accessKeyId = cfg.GetAccessKeyId()
		c.accessKeySecret = cfg.GetAccessKeySecret()
		c.endpoint = cfg.GetEndpoint()
	}
}

// WithLogger sets the logger
func WithLogger(logger *zap.SugaredLogger) ClientOption {
	return func(c *Client) {
		c.logger = logger
	}
}

// initClientV2 initializes the SMS clientV2
func (c *Client) initClientV2() *dysmsapiV2.Client {
	if c.accessKeySecret == "" || c.accessKeyId == "" {
		panic("SMS sending credential information is not configured")
	}
	config := &openapi.Config{
		AccessKeyId:     &c.accessKeyId,
		AccessKeySecret: &c.accessKeySecret,
	}
	config.Endpoint = tea.String(c.endpoint)
	client, err := dysmsapiV2.NewClient(config)
	if err != nil {
		panic(err)
	}
	return client
}

// initClientV3 initializes the SMS clientV3
func (c *Client) initClientV3() *dysmsapiV3.Client {
	if c.accessKeySecret == "" || c.accessKeyId == "" {
		panic("SMS sending credential information is not configured")
	}
	config := &openapi.Config{
		AccessKeyId:     &c.accessKeyId,
		AccessKeySecret: &c.accessKeySecret,
	}
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	client, err := dysmsapiV3.NewClient(config)
	if err != nil {
		panic(err)
	}
	return client
}

var runtimeOptions = &util.RuntimeOptions{
	Autoretry:   tea.Bool(true),
	MaxAttempts: tea.Int(3),
	IgnoreSSL:   tea.Bool(true),
}

// SetRuntimeOptions sets the runtime options
func SetRuntimeOptions(opts ...func(options *util.RuntimeOptions)) func(c *Client) {
	return func(c *Client) {
		for _, opt := range opts {
			opt(runtimeOptions)
		}
	}
}
