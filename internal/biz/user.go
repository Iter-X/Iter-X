package biz

import (
	"context"

	v1 "github.com/iter-x/iter-x/internal/api/user/v1"
	"github.com/iter-x/iter-x/internal/data/ent"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/helper/auth"
)

type User struct {
	cfg *conf.Auth
	repository.Transaction
	userRepo     repository.UserRepo
	languageRepo repository.LanguageRepo
	logger       *zap.SugaredLogger
}

func NewUser(c *conf.Auth, transaction repository.Transaction, userRepo repository.UserRepo, languageRepo repository.LanguageRepo, logger *zap.SugaredLogger) *User {
	return &User{
		cfg:          c,
		Transaction:  transaction,
		userRepo:     userRepo,
		languageRepo: languageRepo,
		logger:       logger.Named("biz.user"),
	}
}

func (b *User) GetUserInfo(ctx context.Context) (*do.User, error) {
	claims, err := auth.ExtractClaims(ctx)
	if err != nil {
		return nil, xerr.ErrorUnauthorized()
	}

	user, err := b.userRepo.FindUserById(ctx, claims.UID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, xerr.ErrorUnauthorized()
		}
		b.logger.Errorw("failed to find user by id", "err", err)
		return nil, xerr.ErrorInternalServerError()
	}

	return user, nil
}

func (b *User) UpdateUserInfo(ctx context.Context, params *bo.UpdateUserInfoRequest) error {
	claims, err := auth.ExtractClaims(ctx)
	if err != nil {
		return xerr.ErrorUnauthorized()
	}

	user, err := b.userRepo.FindUserById(ctx, claims.UID)
	if err != nil {
		if ent.IsNotFound(err) {
			return xerr.ErrorUnauthorized()
		}
		b.logger.Errorw("failed to find user by id", "err", err)
		return xerr.ErrorInternalServerError()
	}

	if params.Username != "" {
		user.Username = params.Username
	}
	if params.Nickname != "" {
		user.Nickname = params.Nickname
	}
	if params.AvatarURL != "" {
		user.AvatarURL = params.AvatarURL
	}

	_, err = b.userRepo.Update(ctx, user)
	return err
}

func (b *User) GetUserPreferences(ctx context.Context) (*do.UserPreference, error) {
	claims, err := auth.ExtractClaims(ctx)
	if err != nil {
		return nil, xerr.ErrorUnauthorized()
	}

	pref, err := b.userRepo.GetUserPreference(ctx, claims.UID)
	if err != nil {
		if ent.IsNotFound(err) {
			return &do.UserPreference{
				AppLanguage:           "en",
				DefaultCity:           "",
				TimeFormat:            "24h",
				DistanceUnit:          "km",
				DarkMode:              "system",
				TripReminder:          true,
				CommunityNotification: true,
				RecommendContentPush:  true,
			}, nil
		}
		return nil, xerr.ErrorInternalServerError()
	}

	return pref, nil
}

func (b *User) UpdateUserPreferences(ctx context.Context, req *bo.UserPreference) error {
	claims, err := auth.ExtractClaims(ctx)
	if err != nil {
		return xerr.ErrorUnauthorized()
	}

	_, err = b.languageRepo.FindLanguageByCode(ctx, req.AppLanguage)
	if err != nil {
		return xerr.ErrorInvalidLanguage()
	}

	pref := &do.UserPreference{
		AppLanguage:           req.AppLanguage,
		DefaultCity:           req.DefaultCity,
		TimeFormat:            req.TimeFormat,
		DistanceUnit:          req.DistanceUnit,
		DarkMode:              req.DarkMode,
		TripReminder:          req.TripReminder,
		CommunityNotification: req.CommunityNotification,
		RecommendContentPush:  req.ContentPush,
	}

	_, err = b.userRepo.GetUserPreference(ctx, claims.UID)
	if err != nil {
		if ent.IsNotFound(err) {
			return b.userRepo.CreateUserPreference(ctx, claims.UID, pref)
		}
		return xerr.ErrorInternalServerError()
	}

	return b.userRepo.UpdateUserPreference(ctx, claims.UID, pref)
}

func (b *User) GetSupportedLanguages(ctx context.Context) ([]*v1.Language, error) {
	languages, err := b.languageRepo.ListLanguages(ctx)
	if err != nil {
		return nil, xerr.ErrorInternalServerError()
	}

	result := make([]*v1.Language, 0, len(languages))
	for _, lang := range languages {
		if !lang.Enabled {
			continue
		}
		result = append(result, &v1.Language{
			Code:       lang.Code,
			Name:       lang.Name,
			NativeName: lang.NativeName,
		})
	}

	return result, nil
}
