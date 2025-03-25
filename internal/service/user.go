package service

import (
	"context"

	userV1 "github.com/iter-x/iter-x/internal/api/user/v1"
	"github.com/iter-x/iter-x/internal/biz"
	"github.com/iter-x/iter-x/internal/biz/bo"
)

type User struct {
	userV1.UnimplementedUserServiceServer
	userBiz *biz.User
}

func NewUser(userBiz *biz.User) *User {
	return &User{
		userBiz: userBiz,
	}
}

func (s *User) GetUserInfo(ctx context.Context, _ *userV1.GetUserInfoRequest) (*userV1.GetUserInfoResponse, error) {
	userInfo, err := s.userBiz.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	return &userV1.GetUserInfoResponse{
		Id:          userInfo.ID.String(),
		Username:    userInfo.Username,
		Nickname:    userInfo.Nickname,
		Email:       userInfo.Email,
		PhoneNumber: userInfo.Phone,
		AvatarUrl:   userInfo.AvatarURL,
		CreatedAt:   userInfo.CreatedAt.Unix(),
		UpdatedAt:   userInfo.UpdatedAt.Unix(),
	}, nil
}

func (s *User) UpdateUserInfo(ctx context.Context, req *userV1.UpdateUserInfoRequest) (*userV1.UpdateUserInfoResponse, error) {
	params := &bo.UpdateUserInfoRequest{
		Username:  req.GetUsername(),
		Nickname:  req.GetNickname(),
		AvatarURL: req.GetAvatarUrl(),
	}
	err := s.userBiz.UpdateUserInfo(ctx, params)
	if err != nil {
		return nil, err
	}
	return &userV1.UpdateUserInfoResponse{}, nil
}
