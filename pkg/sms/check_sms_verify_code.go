package sms

import (
	"context"

	dysmsapiV2 "github.com/alibabacloud-go/dypnsapi-20170525/v2/client"
	"github.com/iter-x/iter-x/pkg/util/pointer"
)

type (
	CheckSmsVerifyCodeParams interface {
		// GetPhoneNumber 手机号
		GetPhoneNumber() string
		// GetVerifyCode 验证码
		GetVerifyCode() string
	}

	CheckSmsVerifyCodeResponse struct {
		*dysmsapiV2.CheckSmsVerifyCodeResponse
	}

	CheckSmsVerifyCodeModel struct {
		// The external ID.
		//
		// example:
		//
		// 1212312
		OutId string `json:"OutId,omitempty" xml:"OutId,omitempty"`
		// The verification results. Valid values:
		//
		// 	- PASS: The verification is successful.
		//
		// 	- UNKNOWN: The verification failed.
		//
		// example:
		//
		// PASS
		VerifyResult string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
	}
)

// GetCheckSmsVerifyCodeModel 获取响应体
func (r *CheckSmsVerifyCodeResponse) GetCheckSmsVerifyCodeModel() *CheckSmsVerifyCodeModel {
	model := pointer.Get(pointer.Get(r.Body).Model)
	return &CheckSmsVerifyCodeModel{
		OutId:        pointer.Get(model.OutId),
		VerifyResult: pointer.Get(model.VerifyResult),
	}
}

// IsOK 响应是否成功
func (r *CheckSmsVerifyCodeResponse) IsOK() bool {
	return pointer.Get(r.StatusCode) == 200 &&
		pointer.Get(pointer.Get(r.Body).Code) == "OK" &&
		pointer.Get(pointer.Get(r.Body).Success) &&
		r.GetCheckSmsVerifyCodeModel().VerifyResult == "PASS"
}

// CheckSmsVerifyCode 验证短信验证码
func (c *Client) CheckSmsVerifyCode(_ context.Context, params CheckSmsVerifyCodeParams) (*CheckSmsVerifyCodeResponse, error) {
	request := &dysmsapiV2.CheckSmsVerifyCodeRequest{
		PhoneNumber: pointer.Of(params.GetPhoneNumber()),
		VerifyCode:  pointer.Of(params.GetVerifyCode()),
	}
	c.logger.Debugw("CheckSmsVerifyCode", "req", request)
	response, err := c.clientV2.CheckSmsVerifyCode(request)
	if err != nil {
		return nil, err
	}
	c.logger.Debugw("CheckSmsVerifyCode", "resp", response)
	return &CheckSmsVerifyCodeResponse{response}, nil
}
