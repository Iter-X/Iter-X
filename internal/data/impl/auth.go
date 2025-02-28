package impl

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/refreshtoken"
	"github.com/iter-x/iter-x/internal/data/ent/user"
)

type Auth struct {
	*data.Tx
	logger *zap.SugaredLogger
}

func NewAuth(tx *data.Tx, logger *zap.SugaredLogger) *Auth {
	return &Auth{
		Tx:     tx,
		logger: logger.Named("repo.auth"),
	}
}

func (r *Auth) FindByEmail(ctx context.Context, email string) (*ent.User, error) {
	cli := r.GetTx(ctx).User

	usr, err := cli.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, xerr.ErrorUserNotFound()
	}
	return usr, err
}

func (r *Auth) FindUserById(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	cli := r.GetTx(ctx).User

	usr, err := cli.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return usr, err
}

func (r *Auth) Create(ctx context.Context, user *ent.User) (*ent.User, error) {
	cli := r.GetTx(ctx).User

	return cli.Create().
		SetUsername(user.Username).SetEmail(user.Email).SetPassword(user.Password).
		Save(ctx)
}

func (r *Auth) Update(ctx context.Context, user *ent.User) (*ent.User, error) {
	cli := r.GetTx(ctx).User

	return cli.UpdateOneID(user.ID).
		SetUsername(user.Username).SetEmail(user.Email).SetAvatarURL(user.AvatarURL).
		Save(ctx)
}

func (r *Auth) GetRefreshTokenByUserId(ctx context.Context, userId uuid.UUID) (*ent.RefreshToken, error) {
	cli := r.GetTx(ctx).RefreshToken

	return cli.Query().Where(refreshtoken.UserID(userId)).Only(ctx)
}

func (r *Auth) GetRefreshToken(ctx context.Context, token string) (*ent.RefreshToken, error) {
	cli := r.GetTx(ctx).RefreshToken

	return cli.Query().Where(refreshtoken.TokenEQ(token)).Only(ctx)
}

func (r *Auth) SaveRefreshToken(ctx context.Context, val *ent.RefreshToken) error {
	cli := r.GetTx(ctx).RefreshToken

	_, err := cli.Create().
		SetToken(val.Token).SetExpiresAt(val.ExpiresAt).SetUserID(val.UserID).
		Save(ctx)
	return err
}

func (r *Auth) UpdateRefreshToken(ctx context.Context, val *ent.RefreshToken) error {
	cli := r.GetTx(ctx).RefreshToken

	_, err := cli.UpdateOneID(val.ID).
		SetToken(val.Token).SetExpiresAt(val.ExpiresAt).SetUpdatedAt(val.UpdatedAt).
		Save(ctx)
	return err
}
