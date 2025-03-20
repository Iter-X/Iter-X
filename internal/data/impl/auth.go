package impl

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/cache"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/refreshtoken"
	"github.com/iter-x/iter-x/internal/data/ent/user"
	"github.com/iter-x/iter-x/internal/data/impl/build"
)

func NewAuth(c *conf.Auth, d *data.Data, logger *zap.SugaredLogger) repository.AuthRepo {
	return &authRepositoryImpl{
		smsConf: c.GetSmsCode(),
		Tx:      d.Tx,
		Cacher:  d.Cache,
		logger:  logger.Named("repo.auth"),
	}
}

type authRepositoryImpl struct {
	smsConf *conf.Auth_SmsCode
	*data.Tx
	cache.Cacher
	logger *zap.SugaredLogger
}

func (r *authRepositoryImpl) ToEntity(po *ent.User) *do.User {
	if po == nil {
		return nil
	}
	return build.AuthRepositoryImplToEntity(po)
}

func (r *authRepositoryImpl) ToEntities(pos []*ent.User) []*do.User {
	if pos == nil {
		return nil
	}
	return build.AuthRepositoryImplToEntities(pos)
}

func (r *authRepositoryImpl) FindByEmail(ctx context.Context, email string) (*do.User, error) {
	cli := r.GetTx(ctx).User

	usr, err := cli.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, xerr.ErrorUserNotFound()
	}
	return r.ToEntity(usr), err
}

func (r *authRepositoryImpl) FindUserById(ctx context.Context, id uuid.UUID) (*do.User, error) {
	cli := r.GetTx(ctx).User
	usr, err := cli.Get(ctx, id)
	return r.ToEntity(usr), err
}

func (r *authRepositoryImpl) Create(ctx context.Context, user *do.User) (*do.User, error) {
	cli := r.GetTx(ctx).User

	row, err := cli.Create().
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetSalt(user.Salt).
		SetNickname(user.Nickname).
		SetRemark(user.Remark).
		SetPhone(user.Phone).
		SetAvatarURL(user.AvatarURL).
		SetEmail(user.Email).
		SetStatus(user.Status.GetValue()).
		Save(ctx)
	return r.ToEntity(row), err
}

func (r *authRepositoryImpl) Update(ctx context.Context, user *do.User) (*do.User, error) {
	cli := r.GetTx(ctx).User

	row, err := cli.UpdateOneID(user.ID).
		SetUsername(user.Username).SetEmail(user.Email).SetAvatarURL(user.AvatarURL).
		Save(ctx)
	return r.ToEntity(row), err
}

func (r *authRepositoryImpl) GetRefreshTokenByUserId(ctx context.Context, userId uuid.UUID) (*do.RefreshToken, error) {
	cli := r.GetTx(ctx).RefreshToken

	row, err := cli.Query().Where(refreshtoken.UserID(userId)).Only(ctx)
	return build.RefreshTokenImplToEntity(row), err
}

func (r *authRepositoryImpl) GetRefreshToken(ctx context.Context, token string) (*do.RefreshToken, error) {
	cli := r.GetTx(ctx).RefreshToken

	row, err := cli.Query().Where(refreshtoken.TokenEQ(token)).Only(ctx)
	return build.RefreshTokenImplToEntity(row), err
}

func (r *authRepositoryImpl) SaveRefreshToken(ctx context.Context, val *do.RefreshToken) error {
	cli := r.GetTx(ctx).RefreshToken

	_, err := cli.Create().
		SetToken(val.Token).SetExpiresAt(val.ExpiresAt).SetUserID(val.UserID).
		Save(ctx)
	return err
}

func (r *authRepositoryImpl) UpdateRefreshToken(ctx context.Context, val *do.RefreshToken) error {
	cli := r.GetTx(ctx).RefreshToken

	_, err := cli.UpdateOneID(val.ID).
		SetToken(val.Token).SetExpiresAt(val.ExpiresAt).SetUpdatedAt(val.UpdatedAt).
		Save(ctx)
	return err
}

func (r *authRepositoryImpl) FindByPhone(ctx context.Context, phone string) (*do.User, error) {
	cli := r.GetTx(ctx).User

	usr, err := cli.Query().Where(user.PhoneEQ(phone)).Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, xerr.ErrorUserNotFound()
	}
	return r.ToEntity(usr), err
}
