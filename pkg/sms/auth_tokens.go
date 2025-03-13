package sms

import (
	"context"

	dysmsapiV2 "github.com/alibabacloud-go/dypnsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/pkg/util/pointer"
	"github.com/iter-x/iter-x/pkg/xcontext"
)

type (
	AuthTokensConfig interface {
		// GetBundleId gets the iOS application ID. Required when OsType is iOS
		GetBundleId() string
		// GetExpireSec gets the token validity period (in seconds), minimum 900, maximum 43200
		GetExpireSec() int64
		// GetPackageName gets the package name. Required when OsType is Android
		GetPackageName() string
		// GetSignName gets the signature. Required when OsType is Android
		GetSignName() string
		// GetSceneCode gets the scheme number
		GetSceneCode() string
		// GetSmsTemplateCode gets the SMS template Code
		GetSmsTemplateCode() string
		// GetSmsCodeExpireSec gets the SMS verification code validity period (in seconds)
		GetSmsCodeExpireSec() int32
	}

	authTokensConfigOptions struct {
		req           *dysmsapiV2.GetSmsAuthTokensRequest
		runtimeOption *util.RuntimeOptions
	}

	AuthTokensConfigOption func(opts *authTokensConfigOptions)

	GetSmsAuthTokensResponse struct {
		*dysmsapiV2.GetSmsAuthTokensResponse
	}

	GetSmsAuthTokensResponseBodyData struct {
		*dysmsapiV2.GetSmsAuthTokensResponseBodyData
	}
)

// IsRequestOK checks if the request is successful
func (r *GetSmsAuthTokensResponse) IsRequestOK() bool {
	if r == nil {
		return false
	}
	return r.GetSmsAuthTokensResponse != nil && pointer.Get(r.StatusCode) == 200 &&
		r.Body != nil && pointer.Get(r.Body.Code) == "OK"
}

// GetBizToken gets the biz token
func (r *GetSmsAuthTokensResponseBodyData) GetBizToken() string {
	if r == nil || r.GetSmsAuthTokensResponseBodyData == nil || r.GetSmsAuthTokensResponseBodyData.BizToken == nil {
		return ""
	}
	return *r.GetSmsAuthTokensResponseBodyData.BizToken
}

// GetExpireTime gets the expire time
func (r *GetSmsAuthTokensResponseBodyData) GetExpireTime() int64 {
	if r == nil || r.GetSmsAuthTokensResponseBodyData == nil || r.GetSmsAuthTokensResponseBodyData.ExpireTime == nil {
		return 0
	}
	return *r.GetSmsAuthTokensResponseBodyData.ExpireTime
}

// GetStsAccessKeyId gets the STS access key ID
func (r *GetSmsAuthTokensResponseBodyData) GetStsAccessKeyId() string {
	if r == nil || r.GetSmsAuthTokensResponseBodyData == nil || r.GetSmsAuthTokensResponseBodyData.StsAccessKeyId == nil {
		return ""
	}
	return *r.GetSmsAuthTokensResponseBodyData.StsAccessKeyId
}

// GetStsAccessKeySecret gets the STS access key secret
func (r *GetSmsAuthTokensResponseBodyData) GetStsAccessKeySecret() string {
	if r == nil || r.GetSmsAuthTokensResponseBodyData == nil || r.GetSmsAuthTokensResponseBodyData.StsAccessKeySecret == nil {
		return ""
	}
	return *r.GetSmsAuthTokensResponseBodyData.StsAccessKeySecret
}

// GetStsToken gets the STS token
func (r *GetSmsAuthTokensResponseBodyData) GetStsToken() string {
	if r == nil || r.GetSmsAuthTokensResponseBodyData == nil || r.GetSmsAuthTokensResponseBodyData.StsToken == nil {
		return ""
	}
	return *r.GetSmsAuthTokensResponseBodyData.StsToken
}

// GetData gets the data from the response
func (r *GetSmsAuthTokensResponse) GetData() *GetSmsAuthTokensResponseBodyData {
	if r == nil || r.GetSmsAuthTokensResponse == nil || r.GetSmsAuthTokensResponse.Body == nil || r.GetSmsAuthTokensResponse.Body.Data == nil {
		return nil
	}
	return &GetSmsAuthTokensResponseBodyData{r.GetSmsAuthTokensResponse.Body.Data}
}

// WithAuthTokensConfigOptionRuntimeOptions sets the runtime options
func WithAuthTokensConfigOptionRuntimeOptions(r *util.RuntimeOptions) AuthTokensConfigOption {
	return func(opts *authTokensConfigOptions) {
		opts.runtimeOption = r
	}
}

// WithAuthTokensConfigOptionOwnerId sets the owner ID
func WithAuthTokensConfigOptionOwnerId(ownerId int64) AuthTokensConfigOption {
	return func(opts *authTokensConfigOptions) {
		opts.req.OwnerId = &ownerId
	}
}

// WithAuthTokensConfigOptionResourceOwnerAccount sets the resource owner account
func WithAuthTokensConfigOptionResourceOwnerAccount(resourceOwnerAccount string) AuthTokensConfigOption {
	return func(opts *authTokensConfigOptions) {
		opts.req.ResourceOwnerAccount = &resourceOwnerAccount
	}
}

// WithAuthTokensConfigOptionResourceOwnerId sets the resource owner ID
func WithAuthTokensConfigOptionResourceOwnerId(resourceOwnerId int64) AuthTokensConfigOption {
	return func(opts *authTokensConfigOptions) {
		opts.req.ResourceOwnerId = &resourceOwnerId
	}
}

// GetSmsAuthTokens gets the SMS authentication tokens
func (c *Client) GetSmsAuthTokens(ctx context.Context, cfg AuthTokensConfig, opts ...AuthTokensConfigOption) (*GetSmsAuthTokensResponse, error) {
	osType, ok := xcontext.ClientTypeFrom(ctx)
	if !ok {
		c.logger.Errorw("Failed to get client type from context")
		return nil, xerr.ErrorInternalServerError()
	}
	getSmsAuthTokensRequest := &dysmsapiV2.GetSmsAuthTokensRequest{
		BundleId:             pointer.Of(cfg.GetBundleId()),
		Expire:               pointer.Of(cfg.GetExpireSec()),
		OsType:               pointer.Of(osType),
		OwnerId:              nil,
		PackageName:          pointer.Of(cfg.GetPackageName()),
		ResourceOwnerAccount: nil,
		ResourceOwnerId:      nil,
		SceneCode:            pointer.Of(cfg.GetSceneCode()),
		SignName:             pointer.Of(cfg.GetSignName()),
		SmsCodeExpire:        pointer.Of(cfg.GetSmsCodeExpireSec()),
		SmsTemplateCode:      pointer.Of(cfg.GetSmsTemplateCode()),
	}
	authOptions := &authTokensConfigOptions{
		req:           getSmsAuthTokensRequest,
		runtimeOption: runtimeOptions,
	}
	for _, opt := range opts {
		opt(authOptions)
	}

	c.logger.Debugw("GetSmsAuthTokens", "req", authOptions)
	getSmsAuthTokensResponse, err := c.clientV2.GetSmsAuthTokensWithOptions(authOptions.req, authOptions.runtimeOption)
	if err != nil {
		c.logger.Errorw("Failed to get SMS authentication tokens", "error", err)
		return nil, err
	}
	c.logger.Debugw("GetSmsAuthTokens", "resp", getSmsAuthTokensResponse)

	return &GetSmsAuthTokensResponse{getSmsAuthTokensResponse}, nil
}

type (
	VerifySmsCodeConfig interface {
		GetPhoneNumber() string
		GetSmsCode() string
		GetSmsToken() string
	}

	verifySmsCodeConfigOptions struct {
		req           *dysmsapiV2.VerifySmsCodeRequest
		runtimeOption *util.RuntimeOptions
	}

	VerifySmsCodeResponse struct {
		*dysmsapiV2.VerifySmsCodeResponse
	}

	VerifySmsCodeConfigOption func(*verifySmsCodeConfigOptions)
)

// IsOK checks if the response is OK
func (r *VerifySmsCodeResponse) IsOK() bool {
	if r == nil {
		return false
	}
	return r.VerifySmsCodeResponse != nil && pointer.Get(r.StatusCode) == 200 &&
		r.Body != nil && pointer.Get(r.Body.Code) == "OK" && pointer.Get(r.Body.Data)
}

// GetCode gets the code
func (r *VerifySmsCodeResponse) GetCode() string {
	if r == nil || r.VerifySmsCodeResponse == nil || r.VerifySmsCodeResponse.Body == nil || r.VerifySmsCodeResponse.Body.Code == nil {
		return ""
	}
	return *r.VerifySmsCodeResponse.Body.Code
}

// GetData gets the data
func (r *VerifySmsCodeResponse) GetData() bool {
	if r == nil || r.VerifySmsCodeResponse == nil || r.VerifySmsCodeResponse.Body == nil || r.VerifySmsCodeResponse.Body.Data == nil {
		return false
	}
	return *r.VerifySmsCodeResponse.Body.Data
}

// GetMessage gets the message
func (r *VerifySmsCodeResponse) GetMessage() string {
	if r == nil || r.VerifySmsCodeResponse == nil || r.VerifySmsCodeResponse.Body == nil || r.VerifySmsCodeResponse.Body.Message == nil {
		return ""
	}
	return *r.VerifySmsCodeResponse.Body.Message
}

// GetRequestId gets the request ID
func (r *VerifySmsCodeResponse) GetRequestId() string {
	if r == nil || r.VerifySmsCodeResponse == nil || r.VerifySmsCodeResponse.Body == nil || r.VerifySmsCodeResponse.Body.RequestId == nil {
		return ""
	}
	return *r.VerifySmsCodeResponse.Body.RequestId
}

// WithVerifySmsCodeConfigOptionRuntimeOptions sets the runtime options
func WithVerifySmsCodeConfigOptionRuntimeOptions(r *util.RuntimeOptions) VerifySmsCodeConfigOption {
	return func(opts *verifySmsCodeConfigOptions) {
		opts.runtimeOption = r
	}
}

// VerifySmsCode verifies the SMS code
func (c *Client) VerifySmsCode(_ context.Context, cfg VerifySmsCodeConfig, opts ...VerifySmsCodeConfigOption) (*VerifySmsCodeResponse, error) {
	verifyCodeRequest := &dysmsapiV2.VerifySmsCodeRequest{
		PhoneNumber: pointer.Of(cfg.GetPhoneNumber()),
		SmsCode:     pointer.Of(cfg.GetSmsCode()),
		SmsToken:    pointer.Of(cfg.GetSmsToken()),
	}
	req := &verifySmsCodeConfigOptions{
		req:           verifyCodeRequest,
		runtimeOption: runtimeOptions,
	}
	for _, opt := range opts {
		opt(req)
	}
	result, err := c.clientV2.VerifySmsCodeWithOptions(req.req, req.runtimeOption)
	if err != nil {
		c.logger.Errorw("Failed to verify SMS code", "error", err)
		return nil, err
	}
	return &VerifySmsCodeResponse{result}, nil
}
