package sms

import (
	"context"
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"

	dysmsapiV3 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	"github.com/iter-x/iter-x/pkg/util/pointer"
)

type (
	SendSmsConfig interface {
		GetSignName() string
		GetTemplateCode() string
		GetTemplateParam() string
		GetPhoneNumbers() string
	}

	SendSmsResponse struct {
		*dysmsapiV3.SendSmsResponse
	}

	SendSmsResponseBody struct {
		*dysmsapiV3.SendSmsResponseBody
	}

	SendSmsConfigOptions struct {
		req           *dysmsapiV3.SendSmsRequest
		runtimeOption *util.RuntimeOptions
	}

	SendSmsConfigOption func(opts *SendSmsConfigOptions)

	QuerySendDetailsConfig interface {
		GetBizId() string
		GetCurrentPage() int64
		GetPageSize() int64
		GetPhoneNumber() string
		GetSendDate() string
	}

	QuerySendDetailsResponse struct {
		*dysmsapiV3.QuerySendDetailsResponse
	}

	QuerySendDetailsResponseBody struct {
		*dysmsapiV3.QuerySendDetailsResponseBody
	}

	QuerySendDetailItem struct {
		Content      string
		ErrCode      string
		OutId        string
		PhoneNum     string
		ReceiveDate  string
		SendDate     string
		SendStatus   int64
		TemplateCode string
	}

	QuerySendDetailsConfigOptions struct {
		req           *dysmsapiV3.QuerySendDetailsRequest
		runtimeOption *util.RuntimeOptions
	}

	QuerySendDetailsConfigOption func(opts *QuerySendDetailsConfigOptions)
)

// IsOK check response is ok
func (r *SendSmsResponse) IsOK() bool {
	if r == nil {
		return false
	}
	return r.SendSmsResponse != nil && pointer.Get(r.StatusCode) == 200 &&
		r.Body != nil && pointer.Get(r.Body.Code) == "OK"
}

// GetBody get body
func (r *SendSmsResponse) GetBody() *SendSmsResponseBody {
	if r == nil || r.SendSmsResponse == nil || r.SendSmsResponse.Body == nil {
		return nil
	}
	return &SendSmsResponseBody{r.SendSmsResponse.Body}
}

// GetBizId get biz id
func (r *SendSmsResponseBody) GetBizId() string {
	if r == nil || r.BizId == nil {
		return ""
	}
	return *r.BizId
}

// GetCode get code
func (r *SendSmsResponseBody) GetCode() string {
	if r == nil || r.Code == nil {
		return ""
	}
	return *r.Code
}

// GetMessage get message
func (r *SendSmsResponseBody) GetMessage() string {
	if r == nil || r.Message == nil {
		return ""
	}
	return *r.Message
}

// GetRequestId get request id
func (r *SendSmsResponseBody) GetRequestId() string {
	if r == nil || r.RequestId == nil {
		return ""
	}
	return *r.RequestId
}

// IsOK check response is ok
func (r *QuerySendDetailsResponse) IsOK() bool {
	if r == nil {
		return false
	}
	return r.QuerySendDetailsResponse != nil && pointer.Get(r.StatusCode) == 200 &&
		r.Body != nil && pointer.Get(r.Body.Code) == "OK"
}

// GetCode get code
func (r *QuerySendDetailsResponseBody) GetCode() string {
	if r == nil || r.Code == nil {
		return ""
	}
	return *r.Code
}

// GetMessage get message
func (r *QuerySendDetailsResponseBody) GetMessage() string {
	if r == nil || r.Message == nil {
		return ""
	}
	return *r.Message
}

// GetRequestId get request id
func (r *QuerySendDetailsResponseBody) GetRequestId() string {
	if r == nil || r.RequestId == nil {
		return ""
	}
	return *r.RequestId
}

// GetTotalCount get total count
func (r *QuerySendDetailsResponseBody) GetTotalCount() int64 {
	if r == nil || r.TotalCount == nil {
		return 0
	}
	total, _ := strconv.ParseInt(*r.TotalCount, 10, 64)
	return total
}

// GetQuerySendDetailItems get query send detail items
func (r *QuerySendDetailsResponseBody) GetQuerySendDetailItems() []*QuerySendDetailItem {
	if r == nil || r.SmsSendDetailDTOs == nil || r.SmsSendDetailDTOs.SmsSendDetailDTO == nil {
		return nil
	}
	items := make([]*QuerySendDetailItem, 0, len(r.SmsSendDetailDTOs.SmsSendDetailDTO))
	for _, item := range r.SmsSendDetailDTOs.SmsSendDetailDTO {
		items = append(items, &QuerySendDetailItem{
			Content:      pointer.Get(item.Content),
			ErrCode:      pointer.Get(item.ErrCode),
			OutId:        pointer.Get(item.OutId),
			PhoneNum:     pointer.Get(item.PhoneNum),
			ReceiveDate:  pointer.Get(item.ReceiveDate),
			SendDate:     pointer.Get(item.SendDate),
			SendStatus:   pointer.Get(item.SendStatus),
			TemplateCode: pointer.Get(item.TemplateCode),
		})
	}
	return items
}

// SendSms send sms
func (c *Client) SendSms(_ context.Context, config SendSmsConfig, opts ...SendSmsConfigOption) (*SendSmsResponse, error) {
	sendSmsRequest := &dysmsapiV3.SendSmsRequest{
		PhoneNumbers:  pointer.Of(config.GetPhoneNumbers()),
		SignName:      pointer.Of(config.GetSignName()),
		TemplateCode:  pointer.Of(config.GetTemplateCode()),
		TemplateParam: pointer.Of(config.GetTemplateParam()),
	}
	req := &SendSmsConfigOptions{
		req:           sendSmsRequest,
		runtimeOption: runtimeOptions,
	}
	for _, opt := range opts {
		opt(req)
	}
	response, err := c.clientV3.SendSmsWithOptions(req.req, req.runtimeOption)
	if err != nil {
		return nil, err
	}
	return &SendSmsResponse{
		response,
	}, nil
}

// QuerySendDetails query send details
func (c *Client) QuerySendDetails(_ context.Context, config QuerySendDetailsConfig, opts ...QuerySendDetailsConfigOption) (*QuerySendDetailsResponse, error) {
	querySendDetailsRequest := &dysmsapiV3.QuerySendDetailsRequest{
		BizId:       pointer.Of(config.GetBizId()),
		CurrentPage: pointer.Of(config.GetCurrentPage()),
		PageSize:    pointer.Of(config.GetPageSize()),
		PhoneNumber: pointer.Of(config.GetPhoneNumber()),
		SendDate:    pointer.Of(config.GetSendDate()),
	}
	req := &QuerySendDetailsConfigOptions{
		req:           querySendDetailsRequest,
		runtimeOption: runtimeOptions,
	}
	for _, opt := range opts {
		opt(req)
	}
	response, err := c.clientV3.QuerySendDetailsWithOptions(req.req, req.runtimeOption)
	if err != nil {
		return nil, err
	}
	return &QuerySendDetailsResponse{
		response,
	}, nil
}

const numberCharset = "0123456789"

// GenerateRandomNumberCode generate random number code
func GenerateRandomNumberCode(length int) string {
	if length <= 0 {
		return ""
	}
	var code strings.Builder
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(numberCharset))))
		if err != nil {
			panic(err)
		}
		code.WriteByte(numberCharset[index.Int64()])
	}
	return code.String()
}
