package biz

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ifuryst/lol"
	"go.uber.org/zap"

	authV1 "github.com/iter-x/iter-x/internal/api/auth/v1"
	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/helper/auth"
	"github.com/iter-x/iter-x/pkg/sms"
	"github.com/iter-x/iter-x/pkg/util/password"
)

type Auth struct {
	cfg *conf.Auth
	repository.Transaction
	authRepo  repository.AuthRepo
	smsClient *sms.Client
	logger    *zap.SugaredLogger
}

func NewAuth(c *conf.Auth, transaction repository.Transaction, authRepo repository.AuthRepo, logger *zap.SugaredLogger) *Auth {
	logger.Debugw("NewAuth", "conf", c)
	smsClient := sms.NewClient(sms.WithClientConfig(c.GetSmsCode()), sms.WithLogger(logger))
	return &Auth{
		cfg:         c,
		Transaction: transaction,
		authRepo:    authRepo,
		smsClient:   smsClient,
		logger:      logger.Named("biz.auth"),
	}
}

func (b *Auth) getToken(ctx context.Context, user *do.User, renew bool) (*bo.SignInResponse, error) {
	token, err := auth.GenerateToken([]byte(b.cfg.Jwt.Secret), auth.Claims{
		UID:       user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Status:    user.Status.IsActive(),
		AvatarURL: user.AvatarURL,
	}, b.cfg.Jwt.Issuer, b.cfg.Jwt.Expiration.AsDuration())
	if err != nil {
		return nil, err
	}

	var refreshToken string
	if renew {
		refreshToken = lol.RandomString(64)
		now := time.Now()
		err = b.Exec(ctx, func(ctx context.Context) error {
			rt, err := b.authRepo.GetRefreshTokenByUserId(ctx, user.ID)
			if err != nil && !ent.IsNotFound(err) {
				return err
			}

			if ent.IsNotFound(err) {
				return b.authRepo.SaveRefreshToken(ctx, &do.RefreshToken{
					Token:     refreshToken,
					ExpiresAt: now.Add(b.cfg.Jwt.RefreshExpiration.AsDuration()),
					CreatedAt: now,
					UserID:    user.ID,
				})
			}

			rt.Token = refreshToken
			rt.ExpiresAt = now.Add(b.cfg.Jwt.RefreshExpiration.AsDuration())
			rt.UpdatedAt = now
			return b.authRepo.UpdateRefreshToken(ctx, rt)
		})
		if err != nil {
			return nil, err
		}
	}

	return &bo.SignInResponse{
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresIn:    b.cfg.Jwt.Expiration.AsDuration().Seconds(),
	}, nil
}

// SignIn signs in and returns a token.
func (b *Auth) SignIn(ctx context.Context, params *authV1.SignInRequest) (*authV1.SignInResponse, error) {
	user, err := b.authRepo.FindByEmail(ctx, params.Email)
	if err != nil || user == nil {
		return nil, xerr.ErrorInvalidEmailOrPassword()
	}

	pass := password.New(params.Password, user.Salt)
	if !pass.EQ(user.Password) {
		return nil, xerr.ErrorInvalidEmailOrPassword()
	}

	res, err := b.getToken(ctx, user, true)
	if err != nil {
		return nil, err
	}
	return &authV1.SignInResponse{
		Token:        res.Token,
		RefreshToken: res.RefreshToken,
		ExpiresIn:    res.ExpiresIn,
	}, nil
}

// SignUp creates a new user.
func (b *Auth) SignUp(ctx context.Context, params *authV1.SignUpRequest) (*authV1.SignUpResponse, error) {
	// confirm the user does not exist
	usr, err := b.authRepo.FindByEmail(ctx, params.Email)
	if err != nil && !xerr.IsUserNotFound(err) {
		return nil, err
	}
	if usr != nil {
		return nil, xerr.ErrorUserAlreadyExists()
	}

	pass := password.New(params.Password)
	passwd, err := pass.EnValue()
	if err != nil {
		return nil, xerr.ErrorInternalServerError()
	}
	// create the user
	data := &do.User{
		Username: params.Email,
		Email:    params.Email,
		Password: passwd,
		Salt:     pass.Salt(),
	}
	user, err := b.authRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	res, err := b.getToken(ctx, user, true)
	if err != nil {
		return nil, err
	}
	return &authV1.SignUpResponse{
		Token:        res.Token,
		RefreshToken: res.RefreshToken,
		ExpiresIn:    res.ExpiresIn,
	}, nil
}

// SignInWithOAuth signs in with oauth and returns a token.
func (b *Auth) SignInWithOAuth(ctx context.Context, params *authV1.SignInWithOAuthRequest) (string, error) {
	var user = &do.User{}

	switch params.Provider {
	case authV1.OAuthProvider_GITHUB:
		res, err := auth.GitHub(ctx, b.cfg.Oauth.GithubClientId,
			b.cfg.Oauth.GithubClientSecret, b.cfg.Oauth.GithubRedirectUrl, params.Code)
		if err != nil {
			return "", err
		}
		user.Username = res["login"].(string)
		user.Email = res["email"].(string)
		user.AvatarURL = res["avatar_url"].(string)
	case authV1.OAuthProvider_GOOGLE:
		res, err := auth.Google(ctx, b.cfg.Oauth.GoogleClientId,
			b.cfg.Oauth.GoogleClientSecret, b.cfg.Oauth.GoogleRedirectUrl, params.Code)
		if err != nil {
			return "", err
		}
		user.Username = res["name"].(string)
		user.Email = res["email"].(string)
		user.AvatarURL = res["picture"].(string)
	case authV1.OAuthProvider_WECHAT:
		res, err := auth.WeChat(ctx, b.cfg.Oauth.WechatClientId,
			b.cfg.Oauth.WechatClientSecret, b.cfg.Oauth.WechatRedirectUrl, params.Code)
		if err != nil {
			return "", err
		}
		user.Username = res["nickname"].(string)
		user.Email = res["email"].(string)
		user.AvatarURL = res["headimgurl"].(string)
	default:
		return "", xerr.ErrorProviderNotSupported()
	}

	usr, err := b.authRepo.FindByEmail(ctx, user.Email)
	if err != nil {
		if !xerr.IsUserNotFound(err) {
			return "", err
		}
		pass := password.New(password.GenerateRandomPassword(8))
		user.Salt = pass.Salt()
		user.Password, err = pass.EnValue()
		if err != nil {
			return "", err
		}
		// create the user
		user, err = b.authRepo.Create(ctx, user)
		if err != nil {
			return "", err
		}
	} else {
		// update information
		usr.Username = user.Username
		usr.AvatarURL = user.AvatarURL
		user, err = b.authRepo.Update(ctx, usr)
		if err != nil {
			return "", err
		}
	}

	return auth.GenerateToken([]byte(b.cfg.Jwt.Secret), auth.Claims{
		UID:       user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Status:    user.Status.IsActive(),
		AvatarURL: user.AvatarURL,
	}, b.cfg.Jwt.Issuer, b.cfg.Jwt.Expiration.AsDuration())
}

func (b *Auth) RefreshToken(ctx context.Context, params *authV1.RefreshTokenRequest) (*authV1.RefreshTokenResponse, error) {
	refreshToken, err := b.authRepo.GetRefreshToken(ctx, params.RefreshToken)
	if err != nil {
		return nil, xerr.ErrorUnauthorized()
	}
	if refreshToken.ExpiresAt.Before(time.Now()) {
		return nil, xerr.ErrorTokenExpired()
	}

	user, err := b.authRepo.FindUserById(ctx, refreshToken.UserID)
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
	return &authV1.RefreshTokenResponse{
		Token:     res.Token,
		ExpiresIn: res.ExpiresIn,
	}, nil
}

func (b *Auth) ValidateToken(_ context.Context, s string) (jwt.Claims, error) {
	token, err := auth.ValidToken(
		[]byte(b.cfg.Jwt.Secret),
		strings.TrimPrefix(s, "Bearer "),
		&auth.Claims{},
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
	claims, ok := token.Claims.(*auth.Claims)
	if !ok {
		return nil, xerr.ErrorInvalidToken()
	}
	return claims, nil
}

// SendSmsCode sends sms code.
func (b *Auth) SendSmsCode(ctx context.Context, params *authV1.SendSmsCodeRequest) (*authV1.SendSmsCodeResponse, error) {
	sendSmsConfigParams := &bo.SendSmsConfigParams{
		PhoneNumber: params.GetPhoneNumber(),
	}
	smsCodeConf := b.cfg.GetSmsCode()
	smsCode := sms.GenerateRandomNumberCode(int(smsCodeConf.GetCodeLength()))
	sendParams := sendSmsConfigParams.WithSmsConfig(smsCodeConf, fmt.Sprintf(`{"code":"%s"}`, smsCode))
	sendSmsResponse, err := b.smsClient.SendSmsVerifyCode(ctx, sendParams)
	if err != nil {
		return nil, err
	}
	if !sendSmsResponse.IsOK() {
		return nil, xerr.ErrorBadRequest()
	}
	return &authV1.SendSmsCodeResponse{
		ExpireTime: smsCodeConf.GetValidTime().GetSeconds(),
		Interval:   smsCodeConf.GetInterval().GetSeconds(),
	}, nil
}

// VerifySmsCode verifies sms code.
func (b *Auth) VerifySmsCode(ctx context.Context, params *authV1.VerifySmsCodeRequest) (*authV1.VerifySmsCodeResponse, error) {
	verifySmsCode, err := b.smsClient.CheckSmsVerifyCode(ctx, params)
	if err != nil {
		return nil, err
	}

	if !verifySmsCode.IsOK() {
		return nil, xerr.ErrorSmsCodeInvalid()
	}

	res, err := b.loginByPhone(ctx, params.GetPhoneNumber())
	if err != nil {
		return nil, err
	}
	return &authV1.VerifySmsCodeResponse{
		Token:     res.Token,
		ExpiresIn: res.ExpiresIn,
	}, nil
}

// OneClickLogin one click login.
func (b *Auth) OneClickLogin(ctx context.Context, params *authV1.OneClickLoginRequest) (*authV1.OneClickLoginResponse, error) {
	if params.GetToken() == "" {
		return nil, xerr.ErrorBadRequest()
	}

	getMobileConfigParams := &bo.GetMobileConfigParams{Token: params.GetToken()}
	// TODO logging request to sms service
	mobileResponse, err := b.smsClient.GetMobile(ctx, getMobileConfigParams)
	if err != nil {
		return nil, err
	}
	if !mobileResponse.IsOK() {
		return nil, xerr.ErrorBadRequest()
	}

	res, err := b.loginByPhone(ctx, mobileResponse.GetBody().GetMobile())
	if err != nil {
		return nil, err
	}
	return &authV1.OneClickLoginResponse{
		Token:     res.Token,
		ExpiresIn: res.ExpiresIn,
	}, nil
}

// loginByPhone logs in by phone.
func (b *Auth) loginByPhone(ctx context.Context, phone string) (*bo.SignInResponse, error) {
	createParams := &bo.CreateUserByPhoneParam{
		PhoneNumber: phone,
	}
	user, err := b.authRepo.FindByPhone(ctx, createParams.PhoneNumber)
	if err != nil {
		if !xerr.IsUserNotFound(err) {
			return nil, err
		}
		user, err = b.authRepo.Create(ctx, createParams.GenerateUserDo())
		if err != nil {
			return nil, err
		}
	}

	return b.getToken(ctx, user, false)
}
