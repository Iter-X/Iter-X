package sms

import (
	"context"

	dysmsapi "github.com/alibabacloud-go/dypnsapi-20170525/v2/client"
	"github.com/iter-x/iter-x/pkg/util/pointer"
)

type (
	GetMobileConfig interface {
		GetAccessToken() string
		GetOutId() string
		GetOwnerId() int64
		GetResourceOwnerAccount() string
		GetResourceOwnerId() int64
	}

	GetMobileResponse struct {
		*dysmsapi.GetMobileResponse
	}

	GetMobileResponseBody struct {
		*dysmsapi.GetMobileResponseBody
	}
)

// GetBody get mobile response body
func (r *GetMobileResponse) GetBody() *GetMobileResponseBody {
	if r == nil || r.GetMobileResponse == nil || r.GetMobileResponse.Body == nil {
		return nil
	}
	return &GetMobileResponseBody{r.GetMobileResponse.Body}
}

// IsOK is ok
func (r *GetMobileResponse) IsOK() bool {
	if r == nil || r.GetMobileResponse == nil || r.GetMobileResponse.Body == nil || r.GetMobileResponse.Body.Code == nil {
		return false
	}
	return *r.GetMobileResponse.Body.Code == "OK"
}

// GetMobile get mobile
func (r *GetMobileResponseBody) GetMobile() string {
	if r == nil || r.GetMobileResponseBody == nil || r.GetMobileResponseBody.GetMobileResultDTO == nil || r.GetMobileResponseBody.GetMobileResultDTO.Mobile == nil {
		return ""
	}
	return *r.GetMobileResponseBody.GetMobileResultDTO.Mobile
}

// GetMobile get mobile
func (c *Client) GetMobile(_ context.Context, cfg GetMobileConfig) (*GetMobileResponse, error) {
	getMobileRequest := &dysmsapi.GetMobileRequest{
		AccessToken:          pointer.Of(cfg.GetAccessToken()),
		OutId:                pointer.Of(cfg.GetOutId()),
		OwnerId:              pointer.Of(cfg.GetOwnerId()),
		ResourceOwnerAccount: pointer.Of(cfg.GetResourceOwnerAccount()),
		ResourceOwnerId:      pointer.Of(cfg.GetResourceOwnerId()),
	}
	getMobileResponse, err := c.client.GetMobile(getMobileRequest)
	if err != nil {
		c.logger.Errorw("Failed to get mobile", "error", err)
		return nil, err
	}
	return &GetMobileResponse{getMobileResponse}, nil
}
