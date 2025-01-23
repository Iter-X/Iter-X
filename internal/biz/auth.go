package biz

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ifuryst/lol"
	v1 "github.com/iter-x/iter-x/internal/api/auth/v1"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/helper/auth"
	"github.com/iter-x/iter-x/internal/repo"
	"github.com/iter-x/iter-x/internal/repo/ent"
	"go.uber.org/zap"
	"strings"
	"time"
)

type (
	Auth struct {
		cfg    *conf.Auth
		repo   *repo.Auth
		logger *zap.SugaredLogger
	}

	SignInResponse struct {
		Token        string
		RefreshToken string
		ExpiresIn    float64
	}
)

func NewAuth(c *conf.Auth, repo *repo.Auth, logger *zap.SugaredLogger) *Auth {
	return &Auth{
		cfg:    c,
		repo:   repo,
		logger: logger.Named("biz.auth"),
	}
}

func (b *Auth) getToken(ctx context.Context, user *ent.User, renew bool) (*SignInResponse, error) {
	token, err := auth.GenerateToken([]byte(b.cfg.Jwt.Secret), auth.Claims{
		UID:       user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Status:    user.Status,
		AvatarURL: user.AvatarURL,
	}, b.cfg.Jwt.Issuer, b.cfg.Jwt.Expiration.AsDuration())
	if err != nil {
		return nil, err
	}

	var refreshToken string
	if renew {
		refreshToken = lol.RandomString(64)
		now := time.Now()
		err = b.repo.Tx.WithTx(ctx, func(tx *ent.Tx) error {
			rt, err := b.repo.GetRefreshTokenByUserId(ctx, user.ID, tx)
			if err != nil && !ent.IsNotFound(err) {
				return err
			}

			if ent.IsNotFound(err) {
				return b.repo.SaveRefreshToken(ctx, &ent.RefreshToken{
					Token:     refreshToken,
					ExpiresAt: now.Add(b.cfg.Jwt.RefreshExpiration.AsDuration()),
					CreatedAt: now,
					UserID:    user.ID,
				}, tx)
			}

			rt.Token = refreshToken
			rt.ExpiresAt = now.Add(b.cfg.Jwt.RefreshExpiration.AsDuration())
			rt.UpdatedAt = now
			return b.repo.UpdateRefreshToken(ctx, rt, tx)
		})
	}

	return &SignInResponse{
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresIn:    b.cfg.Jwt.Expiration.AsDuration().Seconds(),
	}, nil
}

// SignIn signs in and returns a token.
func (b *Auth) SignIn(ctx context.Context, params *v1.SignInRequest) (*v1.SignInResponse, error) {
	user, err := b.repo.FindByEmail(ctx, params.Email)
	if err != nil || user == nil {
		return nil, xerr.ErrorInvalidEmailOrPassword()
	}

	if !auth.CompareHashAndPassword(params.Password, user.Password) {
		return nil, xerr.ErrorInvalidEmailOrPassword()
	}

	res, err := b.getToken(ctx, user, true)
	if err != nil {
		return nil, err
	}
	return &v1.SignInResponse{
		Token:        res.Token,
		RefreshToken: res.RefreshToken,
		ExpiresIn:    res.ExpiresIn,
	}, nil
}

// SignUp creates a new user.
func (b *Auth) SignUp(ctx context.Context, params *v1.SignUpRequest) (*v1.SignUpResponse, error) {
	// confirm the user does not exist
	usr, err := b.repo.FindByEmail(ctx, params.Email)
	if err != nil && !xerr.IsUserNotFound(err) {
		return nil, err
	}
	if usr != nil {
		return nil, xerr.ErrorUserAlreadyExists()
	}

	hashedPass, err := auth.HashPassword(params.Password)
	if err != nil {
		return nil, err
	}

	// create the user
	data := &ent.User{
		Username: params.Email,
		Email:    params.Email,
		Password: hashedPass,
	}
	user, err := b.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	res, err := b.getToken(ctx, user, true)
	if err != nil {
		return nil, err
	}
	return &v1.SignUpResponse{
		Token:        res.Token,
		RefreshToken: res.RefreshToken,
		ExpiresIn:    res.ExpiresIn,
	}, nil
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

func (b *Auth) RefreshToken(ctx context.Context, params *v1.RefreshTokenRequest) (*v1.RefreshTokenResponse, error) {
	refreshToken, err := b.repo.GetRefreshToken(ctx, params.RefreshToken)
	if err != nil {
		return nil, xerr.ErrorUnauthorized()
	}
	if refreshToken.ExpiresAt.Before(time.Now()) {
		return nil, xerr.ErrorTokenExpired()
	}

	user, err := b.repo.FindUserById(ctx, refreshToken.UserID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, xerr.ErrorUnauthorized()
		}
		return nil, err
	}

	res, err := b.getToken(ctx, user, false)
	if err != nil {
		return nil, err
	}
	return &v1.RefreshTokenResponse{
		Token:     res.Token,
		ExpiresIn: res.ExpiresIn,
	}, nil
}

func (b *Auth) ValidateToken(_ context.Context, s string) (jwt.Claims, error) {
	token, err := auth.ValidToken(
		[]byte(b.cfg.Jwt.Secret),
		strings.TrimPrefix(s, "Bearer "),
		&auth.AgentClaims{},
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, xerr.ErrorTokenExpired()
		}
		return nil, xerr.ErrorUnauthorized()
	}
	if !token.Valid {
		return nil, xerr.ErrorInvalidToken()
	}
	claims, ok := token.Claims.(*auth.AgentClaims)
	if !ok {
		return nil, xerr.ErrorInvalidToken()
	}
	return claims, nil
}
