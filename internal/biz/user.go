package biz

import (
	"context"
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
	userRepo repository.UserRepo
	logger   *zap.SugaredLogger
}

func NewUser(c *conf.Auth, transaction repository.Transaction, userRepo repository.UserRepo, logger *zap.SugaredLogger) *User {
	return &User{
		cfg:         c,
		Transaction: transaction,
		userRepo:    userRepo,
		logger:      logger.Named("biz.user"),
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
