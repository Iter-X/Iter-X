package sms

import (
	"context"
	"encoding/json"

	"github.com/iter-x/iter-x/pkg/util/pointer"

	dysmsapiV2 "github.com/alibabacloud-go/dypnsapi-20170525/v2/client"
)

type (
	SendSmsVerifyCodeParams interface {
		// GetPhoneNumber gets the phone number
		GetPhoneNumber() string
		// GetSignName gets the signature name
		GetSignName() string
		// GetTemplateCode gets the template code
		//  SMS template CODE.
		//  You can log in to the SMS service console, select domestic messages or international/Hong Kong, Macao and Taiwan messages, and view the template CODE on the template management page.
		//  https://dysms.console.aliyun.com/dysms.htm?spm=a2c4g.11186623.0.0.7f43422bP2FbbB#/overview
		GetTemplateCode() string
		// GetTemplateParam gets the template parameters
		//  Parameter values for SMS template variables. Use "##code##" for verification code position.
		//  Example: If the template content is: "Your verification code is ${authCode}, valid for 5 minutes, please do not tell others." Then, this field should be: {"authCode":"##code##"}
		//  Note: Replace authCode in the above with the actual parameter name in your verification code template
		GetTemplateParam() string
		// GetValidTime gets the verification code validity period in seconds, default is 300 seconds
		GetValidTime() int64
		// GetCodeLength gets the verification code length, supports 4-8 digits, default is 4 digits
		GetCodeLength() int64
		// GetInterval gets the time interval in seconds between sending verification codes, used for frequency control, default is 60 seconds
		GetInterval() int64
		// GetCodeType gets the generated verification code type
		//   1: Numbers only (default).
		//   2: Uppercase letters only.
		//   3: Lowercase letters only.
		//   4: Mixed case letters.
		//   5: Numbers + uppercase letters.
		//   6: Numbers + lowercase letters.
		//   7: Numbers + mixed case letters.
		GetCodeType() int64
	}

	SendSmsVerifyCodeResponse struct {
		*dysmsapiV2.SendSmsVerifyCodeResponse
	}

	SendSmsVerifyCodeModel struct {
		// The business ID.
		//
		// example:
		//
		// 112231421412414124123^4
		BizId string `json:"BizId,omitempty" xml:"BizId,omitempty"`
		// The external ID.
		//
		// example:
		//
		// 1231231313
		OutId string `json:"OutId,omitempty" xml:"OutId,omitempty"`
		// The request ID.
		//
		// example:
		//
		// API-reqelekrqkllkkewrlwrjlsdfsdf
		RequestId string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
		// The verification code.
		//
		// example:
		//
		// 42324
		VerifyCode string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
	}
)

// String SendSmsVerifyCodeResponse stringifies the response
func (r *SendSmsVerifyCodeResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// IsOK checks if the request status code is OK
func (r *SendSmsVerifyCodeResponse) IsOK() bool {
	return pointer.Get(r.StatusCode) == 200 && pointer.Get(pointer.Get(r.Body).Code) == "OK" && pointer.Get(pointer.Get(r.Body).Success)
}

// GetSendSmsVerifyCodeModel gets the SMS verification code response result
func (r *SendSmsVerifyCodeResponse) GetSendSmsVerifyCodeModel() *SendSmsVerifyCodeModel {
	model := pointer.Get(pointer.Get(r.Body).Model)
	return &SendSmsVerifyCodeModel{
		BizId:      pointer.Get(model.BizId),
		OutId:      pointer.Get(model.OutId),
		RequestId:  pointer.Get(model.RequestId),
		VerifyCode: pointer.Get(model.VerifyCode),
	}
}

// SendSmsVerifyCode
//
//	https://help.aliyun.com/zh/pnvs/developer-reference/api-dypnsapi-2017-05-25-sendsmsverifycode?spm=a2c4g.11186623.help-menu-75010.d_5_5_3_0_0_0.7db318d2Jl66Ok&scm=20140722.H_2573695._.OR_help-T_cn~zh-V_1
func (c *Client) SendSmsVerifyCode(_ context.Context, params SendSmsVerifyCodeParams) (*SendSmsVerifyCodeResponse, error) {
	req := &dysmsapiV2.SendSmsVerifyCodeRequest{
		PhoneNumber:      pointer.Of(params.GetPhoneNumber()),
		SignName:         pointer.Of(params.GetSignName()),
		TemplateCode:     pointer.Of(params.GetTemplateCode()),
		TemplateParam:    pointer.Of(params.GetTemplateParam()),
		CodeLength:       pointer.Of(params.GetCodeLength()),
		ValidTime:        pointer.Of(params.GetValidTime()),
		CodeType:         pointer.Of(params.GetCodeType()),
		Interval:         pointer.Of(params.GetInterval()),
		ReturnVerifyCode: pointer.Of(true),
	}
	c.logger.Debugw("SendSmsVerifyCode", "req", req)
	resp, err := c.clientV2.SendSmsVerifyCode(req)
	if err != nil {
		return nil, err
	}

	res := &SendSmsVerifyCodeResponse{resp}
	c.logger.Debugw("SendSmsVerifyCode", "resp", res, "code", res.GetSendSmsVerifyCodeModel().VerifyCode)
	return res, nil
}
