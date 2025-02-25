package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/repo/ent"
	"github.com/iter-x/iter-x/internal/repo/ent/refreshtoken"
	"github.com/iter-x/iter-x/internal/repo/ent/user"
	"go.uber.org/zap"
)

type Auth struct {
	*Tx
	cli    *ent.Client
	logger *zap.SugaredLogger
}

func NewAuth(cli *ent.Client, logger *zap.SugaredLogger) *Auth {
	return &Auth{
		Tx:     &Tx{cli: cli},
		cli:    cli,
		logger: logger.Named("repo.auth"),
	}
}

func (r *Auth) FindByEmail(ctx context.Context, email string, tx ...*ent.Tx) (*ent.User, error) {
	cli := r.cli.User
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].User
	}

	usr, err := cli.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, xerr.ErrorUserNotFound()
	}
	return usr, err
}

func (r *Auth) FindUserById(ctx context.Context, id uuid.UUID, tx ...*ent.Tx) (*ent.User, error) {
	cli := r.cli.User
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].User
	}

	usr, err := cli.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return usr, err
}

func (r *Auth) Create(ctx context.Context, user *ent.User, tx ...*ent.Tx) (*ent.User, error) {
	cli := r.cli.User
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].User
	}

	return cli.Create().
		SetUsername(user.Username).SetEmail(user.Email).SetPassword(user.Password).
		Save(ctx)
}

func (r *Auth) Update(ctx context.Context, user *ent.User, tx ...*ent.Tx) (*ent.User, error) {
	cli := r.cli.User
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].User
	}

	return cli.UpdateOneID(user.ID).
		SetUsername(user.Username).SetEmail(user.Email).SetAvatarURL(user.AvatarURL).
		Save(ctx)
}

func (r *Auth) GetRefreshTokenByUserId(ctx context.Context, userId uuid.UUID, tx ...*ent.Tx) (*ent.RefreshToken, error) {
	cli := r.cli.RefreshToken
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].RefreshToken
	}

	return cli.Query().Where(refreshtoken.UserID(userId)).Only(ctx)
}

func (r *Auth) GetRefreshToken(ctx context.Context, token string, tx ...*ent.Tx) (*ent.RefreshToken, error) {
	cli := r.cli.RefreshToken
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].RefreshToken
	}

	return cli.Query().Where(refreshtoken.TokenEQ(token)).Only(ctx)
}

func (r *Auth) SaveRefreshToken(ctx context.Context, val *ent.RefreshToken, tx ...*ent.Tx) error {
	cli := r.cli.RefreshToken
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].RefreshToken
	}

	_, err := cli.Create().
		SetToken(val.Token).SetExpiresAt(val.ExpiresAt).SetUserID(val.UserID).
		Save(ctx)
	return err
}

func (r *Auth) UpdateRefreshToken(ctx context.Context, val *ent.RefreshToken, tx ...*ent.Tx) error {
	cli := r.cli.RefreshToken
	if len(tx) > 0 && tx[0] != nil {
		cli = tx[0].RefreshToken
	}

	_, err := cli.UpdateOneID(val.ID).
		SetToken(val.Token).SetExpiresAt(val.ExpiresAt).SetUpdatedAt(val.UpdatedAt).
		Save(ctx)
	return err
}
