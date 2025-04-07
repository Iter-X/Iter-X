package service

import (
	"context"
	"github.com/iter-x/iter-x/internal/service/build"

	"github.com/iter-x/iter-x/internal/common/xerr"

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

func (s *User) GetUserPreferences(ctx context.Context, _ *userV1.GetUserPreferencesRequest) (*userV1.GetUserPreferencesResponse, error) {
	preferences, err := s.userBiz.GetUserPreferences(ctx)
	if err != nil {
		return nil, err
	}

	return &userV1.GetUserPreferencesResponse{
		Preference: &userV1.UserPreference{
			AppLanguage:           preferences.AppLanguage,
			DefaultCity:           preferences.DefaultCity,
			TimeFormat:            build.GetTimeFormatProto(preferences.TimeFormat),
			DistanceUnit:          build.GetDistanceUnitProto(preferences.DistanceUnit),
			DarkMode:              build.GetDarkModeProto(preferences.DarkMode),
			TripReminder:          preferences.TripReminder,
			CommunityNotification: preferences.CommunityNotification,
			ContentPush:           preferences.RecommendContentPush,
		},
	}, nil
}

func (s *User) UpdateUserPreferences(ctx context.Context, req *userV1.UpdateUserPreferencesRequest) (*userV1.UpdateUserPreferencesResponse, error) {
	pref := req.Preference
	if pref == nil {
		return nil, xerr.ErrorBadRequest()
	}

	err := s.userBiz.UpdateUserPreferences(ctx, &bo.UserPreference{
		AppLanguage:           pref.AppLanguage,
		DefaultCity:           pref.DefaultCity,
		TimeFormat:            build.GetTimeFormatString(pref.TimeFormat),
		DistanceUnit:          build.GetDistanceUnitString(pref.DistanceUnit),
		DarkMode:              build.GetDarkModeString(pref.DarkMode),
		TripReminder:          pref.TripReminder,
		CommunityNotification: pref.CommunityNotification,
		ContentPush:           pref.ContentPush,
	})
	if err != nil {
		return nil, err
	}

	return &userV1.UpdateUserPreferencesResponse{}, nil
}

func (s *User) GetSupportedLanguages(ctx context.Context, _ *userV1.GetSupportedLanguagesRequest) (*userV1.GetSupportedLanguagesResponse, error) {
	languages, err := s.userBiz.GetSupportedLanguages(ctx)
	if err != nil {
		return nil, err
	}

	return &userV1.GetSupportedLanguagesResponse{
		Languages: languages,
	}, nil
}
