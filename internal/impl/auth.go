package impl

import (
	"context"

	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/impl/ent"
	"github.com/iter-x/iter-x/internal/impl/ent/refreshtoken"
	"github.com/iter-x/iter-x/internal/impl/ent/user"
	"go.uber.org/zap"
)

func NewAuthRepository(cli *ent.Client, logger *zap.SugaredLogger) repository.Auth[*ent.User, *do.User] {
	return &authRepositoryImpl{
		Tx:                     &Tx{cli: cli},
		logger:                 logger.Named("repo.auth"),
		refreshTokenEntityImpl: new(refreshTokenRepositoryImpl),
	}
}

type authRepositoryImpl struct {
	*Tx
	logger *zap.SugaredLogger

	refreshTokenEntityImpl repository.Base[*ent.RefreshToken, *do.RefreshToken]
}

func (r *authRepositoryImpl) ToEntity(po *ent.User) *do.User {
	if po == nil {
		return nil
	}
	return &do.User{
		ID:            po.ID,
		CreatedAt:     po.CreatedAt,
		UpdatedAt:     po.UpdatedAt,
		Status:        po.Status,
		Username:      po.Username,
		Password:      po.Password,
		Email:         po.Email,
		AvatarURL:     po.AvatarURL,
		Trips:         nil,
		RefreshTokens: r.refreshTokenEntityImpl.ToEntities(po.Edges.RefreshToken),
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

	usr, err := cli.Create().
		SetUsername(user.Username).SetEmail(user.Email).SetPassword(user.Password).
		Save(ctx)
	return r.ToEntity(usr), err
}

func (r *authRepositoryImpl) Update(ctx context.Context, user *do.User) (*do.User, error) {
	cli := r.GetTx(ctx).User
	usr, err := cli.UpdateOneID(user.ID).
		SetUsername(user.Username).SetEmail(user.Email).SetAvatarURL(user.AvatarURL).
		Save(ctx)
	return r.ToEntity(usr), err
}

func (r *authRepositoryImpl) GetRefreshTokenByUserId(ctx context.Context, userId uuid.UUID) (*do.RefreshToken, error) {
	cli := r.GetTx(ctx).RefreshToken
	token, err := cli.Query().Where(refreshtoken.UserID(userId)).Only(ctx)
	return r.refreshTokenEntityImpl.ToEntity(token), err
}

func (r *authRepositoryImpl) GetRefreshToken(ctx context.Context, token string) (*do.RefreshToken, error) {
	cli := r.GetTx(ctx).RefreshToken

	resToken, err := cli.Query().Where(refreshtoken.TokenEQ(token)).Only(ctx)
	return r.refreshTokenEntityImpl.ToEntity(resToken), err
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
