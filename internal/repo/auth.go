package repo

import (
	"context"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/repo/ent"
	"github.com/iter-x/iter-x/internal/repo/ent/user"
	"go.uber.org/zap"
)

type Auth struct {
	cli    *ent.Client
	logger *zap.SugaredLogger
}

func NewAuth(client *ent.Client, logger *zap.SugaredLogger) *Auth {
	return &Auth{
		cli:    client,
		logger: logger.Named("repo.auth"),
	}
}

func (r *Auth) FindByEmail(ctx context.Context, email string) (*ent.User, error) {
	usr, err := r.cli.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, xerr.ErrorUserNotFound()
	}
	return usr, err
}

func (r *Auth) Create(ctx context.Context, user *ent.User) (*ent.User, error) {
	return r.cli.User.Create().
		SetUsername(user.Username).SetEmail(user.Email).SetPassword(user.Password).
		Save(ctx)
}

func (r *Auth) Update(ctx context.Context, user *ent.User) (*ent.User, error) {
	return r.cli.User.UpdateOneID(user.ID).
		SetUsername(user.Username).SetEmail(user.Email).SetAvatarURL(user.AvatarURL).
		Save(ctx)
}
