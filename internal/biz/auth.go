package biz

import (
	"context"
	"github.com/google/uuid"
	v1 "github.com/iter-x/iter-x/internal/api/auth/v1"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/helper/auth"
	"github.com/iter-x/iter-x/internal/repo"
	"github.com/iter-x/iter-x/internal/repo/ent"
	"go.uber.org/zap"
)

type Auth struct {
	cfg    *conf.Auth
	repo   *repo.Auth
	logger *zap.SugaredLogger
}

func NewAuth(c *conf.Auth, repo *repo.Auth, logger *zap.SugaredLogger) *Auth {
	return &Auth{
		cfg:    c,
		repo:   repo,
		logger: logger.Named("biz.auth"),
	}
}

// SignIn signs in and returns a token.
func (b *Auth) SignIn(ctx context.Context, params *v1.SignInRequest) (string, error) {
	usr, err := b.repo.FindByEmail(ctx, params.Email)
	if err != nil || usr == nil {
		return "", xerr.ErrorInvalidEmailOrPassword()
	}

	if !auth.CompareHashAndPassword(params.Password, usr.Password) {
		return "", xerr.ErrorInvalidEmailOrPassword()
	}

	return auth.GenerateToken([]byte(b.cfg.Jwt.Secret), auth.Claims{
		UID:       usr.ID,
		Username:  usr.Username,
		Email:     usr.Email,
		Status:    usr.Status,
		AvatarURL: usr.AvatarURL,
	}, b.cfg.Jwt.Issuer, b.cfg.Jwt.Expiration.AsDuration())
}

// SignUp creates a new user.
func (b *Auth) SignUp(ctx context.Context, params *v1.SignUpRequest) (string, error) {
	// confirm the user does not exist
	usr, err := b.repo.FindByEmail(ctx, params.Email)
	if err != nil && !xerr.IsUserNotFound(err) {
		return "", err
	}
	if usr != nil {
		return "", xerr.ErrorUserAlreadyExists()
	}

	hashedPass, err := auth.HashPassword(params.Password)
	if err != nil {
		return "", err
	}

	// create the user
	data := &ent.User{
		Username: params.Email,
		Email:    params.Email,
		Password: hashedPass,
	}
	user, err := b.repo.Create(ctx, data)
	if err != nil {
		return "", err
	}

	// generate the token
	return auth.GenerateToken([]byte(b.cfg.Jwt.Secret), auth.Claims{
		UID:       user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Status:    user.Status,
		AvatarURL: user.AvatarURL,
	}, b.cfg.Jwt.Issuer, b.cfg.Jwt.Expiration.AsDuration())
}

// SignInWithOAuth signs in with oauth and returns a token.
func (b *Auth) SignInWithOAuth(ctx context.Context, params *v1.SignInWithOAuthRequest) (string, error) {
	var user = &ent.User{}

	switch params.Provider {
	case v1.OAuthProvider_GITHUB:
		res, err := auth.GitHub(ctx, b.cfg.Oauth.GithubClientId,
			b.cfg.Oauth.GithubClientSecret, b.cfg.Oauth.GithubRedirectUrl, params.Code)
		if err != nil {
			return "", err
		}
		user.Username = res["login"].(string)
		user.Email = res["email"].(string)
		user.AvatarURL = res["avatar_url"].(string)
	case v1.OAuthProvider_GOOGLE:
		res, err := auth.Google(ctx, b.cfg.Oauth.GoogleClientId,
			b.cfg.Oauth.GoogleClientSecret, b.cfg.Oauth.GoogleRedirectUrl, params.Code)
		if err != nil {
			return "", err
		}
		user.Username = res["name"].(string)
		user.Email = res["email"].(string)
		user.AvatarURL = res["picture"].(string)
	default:
		return "", xerr.ErrorProviderNotSupported()
	}

	usr, err := b.repo.FindByEmail(ctx, user.Email)
	if err != nil {
		if !xerr.IsUserNotFound(err) {
			return "", err
		}
		user.Password, err = auth.HashPassword(uuid.New().String())
		if err != nil {
			return "", err
		}
		// create the user
		user, err = b.repo.Create(ctx, user)
		if err != nil {
			return "", err
		}
	} else {
		// update information
		usr.Username = user.Username
		usr.AvatarURL = user.AvatarURL
		user, err = b.repo.Update(ctx, usr)
		if err != nil {
			return "", err
		}
	}

	return auth.GenerateToken([]byte(b.cfg.Jwt.Secret), auth.Claims{
		UID:       user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Status:    user.Status,
		AvatarURL: user.AvatarURL,
	}, b.cfg.Jwt.Issuer, b.cfg.Jwt.Expiration.AsDuration())
}
