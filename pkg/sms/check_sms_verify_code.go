package sms

import (
	"context"
	"encoding/json"

	dysmsapiV2 "github.com/alibabacloud-go/dypnsapi-20170525/v2/client"

	"github.com/iter-x/iter-x/pkg/util/pointer"
)

type (
	CheckSmsVerifyCodeParams interface {
		// GetPhoneNumber gets the phone number
		GetPhoneNumber() string
		// GetVerifyCode gets the verification code
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

// String returns the verification result
func (r *CheckSmsVerifyCodeResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// GetCheckSmsVerifyCodeModel gets the response body
func (r *CheckSmsVerifyCodeResponse) GetCheckSmsVerifyCodeModel() *CheckSmsVerifyCodeModel {
	model := pointer.Get(pointer.Get(r.Body).Model)
	return &CheckSmsVerifyCodeModel{
		OutId:        pointer.Get(model.OutId),
		VerifyResult: pointer.Get(model.VerifyResult),
	}
}

// IsOK checks if the response is successful
func (r *CheckSmsVerifyCodeResponse) IsOK() bool {
	return pointer.Get(r.StatusCode) == 200 &&
		pointer.Get(pointer.Get(r.Body).Code) == "OK" &&
		pointer.Get(pointer.Get(r.Body).Success) &&
		r.GetCheckSmsVerifyCodeModel().VerifyResult == "PASS"
}

// CheckSmsVerifyCode verifies the SMS verification code
func (c *Client) CheckSmsVerifyCode(_ context.Context, params CheckSmsVerifyCodeParams) (*CheckSmsVerifyCodeResponse, error) {
	request := &dysmsapiV2.CheckSmsVerifyCodeRequest{
		PhoneNumber: pointer.Of(params.GetPhoneNumber()),
		VerifyCode:  pointer.Of(params.GetVerifyCode()),
	}
	c.logger.Debugw("CheckSmsVerifyCode", "req", request)
	response, err := c.clientV2.CheckSmsVerifyCode(request)
	if err != nil {
		c.logger.Errorw("Failed to check sms verify code", "error", err)
		return nil, err
	}
	res := &CheckSmsVerifyCodeResponse{response}
	c.logger.Debugw("CheckSmsVerifyCode", "resp", res)
	return res, nil
}
