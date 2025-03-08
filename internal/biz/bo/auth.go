package bo

import (
	"encoding"
	"encoding/json"
	"time"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/pkg/sms"
	"github.com/iter-x/iter-x/pkg/util/password"
	"github.com/iter-x/iter-x/pkg/vobj"
)

var _ encoding.BinaryMarshaler = (*SmsCodeItem)(nil)
var _ encoding.BinaryUnmarshaler = (*SmsCodeItem)(nil)

type (
	SignInResponse struct {
		Token        string
		RefreshToken string
		ExpiresIn    float64
	}

	CreateUserByPhoneParam struct {
		PhoneNumber string
	}

	GetMobileConfigParams struct {
		Token string
	}

	SendSmsConfigParams struct {
		PhoneNumber string
	}

	SmsCodeItem struct {
		PhoneNumber string
		BizToken    string
		Expire      time.Duration
		SmsCode     string
	}

	sendSmsConfigParams struct {
		*SendSmsConfigParams
		signName      string
		templateCode  string
		templateParam string
		validTime     int64
		codeLength    int64
		interval      int64
		codeType      int64
	}
)

// MarshalBinary marshal binary
func (s *SmsCodeItem) MarshalBinary() (data []byte, err error) {
	return json.Marshal(s)
}

// UnmarshalBinary unmarshal binary
func (s *SmsCodeItem) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

func (s *SendSmsConfigParams) WithSmsConfig(smsCodeConf *conf.Auth_SmsCode, templateParam string) sms.SendSmsVerifyCodeParams {
	return &sendSmsConfigParams{
		SendSmsConfigParams: s,
		signName:            smsCodeConf.GetSignName(),
		templateCode:        smsCodeConf.GetTemplateCode(),
		templateParam:       templateParam,
		validTime:           smsCodeConf.GetValidTime().GetSeconds(),
		codeLength:          smsCodeConf.GetCodeLength(),
		interval:            smsCodeConf.GetInterval().GetSeconds(),
		codeType:            1,
	}
}

func (s *sendSmsConfigParams) GetValidTime() int64 {
	return s.validTime
}

func (s *sendSmsConfigParams) GetCodeLength() int64 {
	return s.codeLength
}

func (s *sendSmsConfigParams) GetInterval() int64 {
	return s.interval
}

func (s *sendSmsConfigParams) GetCodeType() int64 {
	return s.codeType
}

func (s *sendSmsConfigParams) GetPhoneNumber() string {
	return s.PhoneNumber
}

func (s *sendSmsConfigParams) GetSignName() string {
	return s.signName
}

func (s *sendSmsConfigParams) GetTemplateCode() string {
	return s.templateCode
}

func (s *sendSmsConfigParams) GetTemplateParam() string {
	return s.templateParam
}

func (g *GetMobileConfigParams) GetAccessToken() string {
	return g.Token
}

func (g *GetMobileConfigParams) GetOutId() string {
	// TODO generate outId
	return ""
}

func (g *GetMobileConfigParams) GetOwnerId() int64 {
	return 0
}

func (g *GetMobileConfigParams) GetResourceOwnerAccount() string {
	// TODO generate resourceOwnerAccount
	return ""
}

func (g *GetMobileConfigParams) GetResourceOwnerId() int64 {
	return 0
}

// GenerateUserDo generate user do
func (c *CreateUserByPhoneParam) GenerateUserDo() *do.User {
	passwordObj := password.New(password.GenerateRandomPassword(8))
	enPass, err := passwordObj.EnValue()
	if err != nil {
		panic(err)
	}
	return &do.User{
		Status:   vobj.UserStatusActive,
		Username: c.PhoneNumber,
		Password: enPass,
		Salt:     passwordObj.Salt(),
		Phone:    c.PhoneNumber,
	}
}
