package sms

import (
	"context"

	"github.com/iter-x/iter-x/pkg/util/pointer"

	dysmsapiV2 "github.com/alibabacloud-go/dypnsapi-20170525/v2/client"
)

type (
	SendSmsVerifyCodeParams interface {
		// GetPhoneNumber 手机号
		GetPhoneNumber() string
		// GetSignName 签名名称
		GetSignName() string
		// GetTemplateCode 模板code
		//  短信模板 CODE。
		//  您可以登录短信服务控制台，选择国内消息或国际/港澳台消息，在模板管理页面查看模板 CODE。
		//  https://dysms.console.aliyun.com/dysms.htm?spm=a2c4g.11186623.0.0.7f43422bP2FbbB#/overview
		GetTemplateCode() string
		// GetTemplateParam 模板参数
		//
		//  短信模板变量填写的参数值。验证码位置使用"##code##"替代。
		//  示例：如模板内容为：“您的验证码是${authCode}，5 分钟内有效，请勿告诉他人。”。此时，该字段传入：{"authCode":"##code##"}
		//  注意 上文中的 authCode 请替换成您实际申请的验证码模板中的参数名称
		GetTemplateParam() string
		// GetValidTime 验证码有效时长，单位秒，默认为 300 秒。
		GetValidTime() int64
		// GetCodeLength 验证码长度支持 4～8 位长度，默认是 4 位。
		GetCodeLength() int64
		// GetInterval 时间间隔，单位：秒。即多久间隔可以发送一次验证码，用于频控，默认 60 秒。
		GetInterval() int64
		// GetCodeType 生成的验证码类型。取值：
		//
		//   1：纯数字（默认）。
		//   2：纯大写字母。
		//   3：纯小写字母。
		//   4：大小字母混合。
		//   5：数字+大写字母混合。
		//   6：数字+小写字母混合。
		//   7：数字+大小写字母混合。
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

// IsOK 请求状态码。返回 OK 代表请求成功。其他错误码，请参见返回码列表。
func (r *SendSmsVerifyCodeResponse) IsOK() bool {
	return pointer.Get(r.StatusCode) == 200 && pointer.Get(pointer.Get(r.Body).Code) == "OK" && pointer.Get(pointer.Get(r.Body).Success)
}

// GetSendSmsVerifyCodeModel 获取短信验证码返回结果
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
		PhoneNumber:   pointer.Of(params.GetPhoneNumber()),
		SignName:      pointer.Of(params.GetSignName()),
		TemplateCode:  pointer.Of(params.GetTemplateCode()),
		TemplateParam: pointer.Of(params.GetTemplateParam()),
		CodeLength:    pointer.Of(params.GetCodeLength()),
		ValidTime:     pointer.Of(params.GetValidTime()),
		CodeType:      pointer.Of(params.GetCodeType()),
		Interval:      pointer.Of(params.GetInterval()),
	}
	resp, err := c.clientV2.SendSmsVerifyCode(req)
	if err != nil {
		return nil, err
	}
	return &SendSmsVerifyCodeResponse{resp}, nil
}
