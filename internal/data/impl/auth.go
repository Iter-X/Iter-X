package impl

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/data/cache"
	"github.com/iter-x/iter-x/pkg/sms"
	"github.com/iter-x/iter-x/pkg/util/password"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/refreshtoken"
	"github.com/iter-x/iter-x/internal/data/ent/user"
	"github.com/iter-x/iter-x/pkg/vobj"
)

func NewAuth(c *conf.Auth, d *data.Data, logger *zap.SugaredLogger) repository.AuthRepo {
	return &authRepositoryImpl{
		smsConf:            c.GetSmsCode(),
		Tx:                 d.Tx,
		Cacher:             d.Cache,
		logger:             logger.Named("repo.auth"),
		refreshTokenImpl:   new(refreshTokenImpl),
		tripRepositoryImpl: new(tripRepositoryImpl),
	}
}

type authRepositoryImpl struct {
	smsConf *conf.Auth_SmsCode
	*data.Tx
	cache.Cacher
	logger *zap.SugaredLogger

	refreshTokenImpl   repository.BaseRepo[*ent.RefreshToken, *do.RefreshToken]
	tripRepositoryImpl repository.BaseRepo[*ent.Trip, *do.Trip]
}

func (r *authRepositoryImpl) ToEntity(po *ent.User) *do.User {
	if po == nil {
		return nil
	}
	return &do.User{
		ID:            po.ID,
		CreatedAt:     po.CreatedAt,
		UpdatedAt:     po.UpdatedAt,
		Status:        vobj.UserStatus(po.Status),
		Username:      po.Username,
		Password:      po.Password,
		Salt:          po.Salt,
		Nickname:      po.Nickname,
		Remark:        po.Remark,
		Phone:         po.Phone,
		Email:         po.Email,
		AvatarURL:     po.AvatarURL,
		RefreshTokens: r.refreshTokenImpl.ToEntities(po.Edges.RefreshToken),
		Trips:         r.tripRepositoryImpl.ToEntities(po.Edges.Trip),
	}
}

func (r *authRepositoryImpl) ToEntities(pos []*ent.User) []*do.User {
	if pos == nil {
		return nil
	}
	list := make([]*do.User, 0, len(pos))
	for _, v := range pos {
		list = append(list, r.ToEntity(v))
	}
	return list
}

func (r *authRepositoryImpl) GetSmsCode(ctx context.Context, params *bo.SendSmsConfigParams) (*bo.SmsCodeItem, error) {
	// generate sms code
	smsCode := sms.GenerateRandomNumberCode(6)
	hash := password.MD5(params.ClientToken)
	expire := r.smsConf.GetExpire().AsDuration()
	bizToken := password.GenerateRandomPassword(32)
	cacheResponse := &bo.SmsCodeItem{
		PhoneNumber: params.PhoneNumber,
		ClientToken: params.ClientToken,
		BizToken:    bizToken,
		Expire:      expire,
		SmsCode:     smsCode,
	}
	// cache sms code
	if err := r.Client().Set(ctx, cache.SmsCodeKey.Key(params.PhoneNumber, "md5", hash), cacheResponse, expire).Err(); err != nil {
		return nil, err
	}
	return cacheResponse, nil
}

func (r *authRepositoryImpl) VerifySmsCode(ctx context.Context, params *bo.VerifySmsCodeParams) error {
	hash := password.MD5(params.ClientToken)
	var cacheResponse bo.SmsCodeItem
	if err := r.Client().Get(ctx, cache.SmsCodeKey.Key(params.PhoneNumber, "md5", hash)).Scan(&cacheResponse); err != nil {
		if errors.Is(err, redis.Nil) {
			return xerr.ErrorSmsCodeExpired()
		}
		return err
	}
	if err := params.Validate(&cacheResponse); err != nil {
		return xerr.ErrorSmsCodeInvalid()
	}
	return nil
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
	return r.refreshTokenImpl.ToEntity(row), err
}

func (r *authRepositoryImpl) GetRefreshToken(ctx context.Context, token string) (*do.RefreshToken, error) {
	cli := r.GetTx(ctx).RefreshToken

	row, err := cli.Query().Where(refreshtoken.TokenEQ(token)).Only(ctx)
	return r.refreshTokenImpl.ToEntity(row), err
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
