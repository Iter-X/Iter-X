package service

import (
	"context"

	"github.com/golang-jwt/jwt/v5"

	authV1 "github.com/iter-x/iter-x/internal/api/auth/v1"
	"github.com/iter-x/iter-x/internal/biz"
	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/pkg/vobj"
)

type Auth struct {
	authV1.UnimplementedAuthServiceServer
	authBiz *biz.Auth
}

func NewAuth(authBiz *biz.Auth) *Auth {
	return &Auth{
		authBiz: authBiz,
	}
}

func (s *Auth) SignIn(ctx context.Context, req *authV1.SignInRequest) (*authV1.SignInResponse, error) {
	params := &bo.SignInRequest{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	signIn, err := s.authBiz.SignIn(ctx, params)
	if err != nil {
		return nil, err
	}
	return &authV1.SignInResponse{
		Token:        signIn.Token,
		RefreshToken: signIn.RefreshToken,
		ExpiresIn:    signIn.ExpiresIn,
	}, nil
}

func (s *Auth) SignInWithOAuth(ctx context.Context, req *authV1.SignInWithOAuthRequest) (*authV1.SignInWithOAuthResponse, error) {
	params := &bo.SignInWithOAuthRequest{
		Code:     req.GetCode(),
		Provider: vobj.OAuthProvider(req.GetProvider()),
	}
	token, err := s.authBiz.SignInWithOAuth(ctx, params)
	if err != nil {
		return nil, err
	}
	return &authV1.SignInWithOAuthResponse{Token: token}, nil
}

func (s *Auth) SignUp(ctx context.Context, req *authV1.SignUpRequest) (*authV1.SignUpResponse, error) {
	params := &bo.SignInRequest{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	signUp, err := s.authBiz.SignUp(ctx, params)
	if err != nil {
		return nil, err
	}
	return &authV1.SignUpResponse{
		Token:        signUp.Token,
		RefreshToken: signUp.RefreshToken,
		ExpiresIn:    signUp.ExpiresIn,
	}, nil
}

func (s *Auth) RequestPasswordReset(ctx context.Context, req *authV1.RequestPasswordResetRequest) (*authV1.RequestPasswordResetResponse, error) {
	// TODO: Implement logic for requesting a password reset
	return &authV1.RequestPasswordResetResponse{
		Status: "password reset request sent",
	}, nil
}

func (s *Auth) VerifyPasswordResetToken(ctx context.Context, req *authV1.VerifyPasswordResetTokenRequest) (*authV1.VerifyPasswordResetTokenResponse, error) {
	// TODO: Implement logic for verifying the password reset token
	return &authV1.VerifyPasswordResetTokenResponse{
		Valid: true,
	}, nil
}

func (s *Auth) ResetPassword(ctx context.Context, req *authV1.ResetPasswordRequest) (*authV1.ResetPasswordResponse, error) {
	// TODO: Implement logic for resetting the password
	return &authV1.ResetPasswordResponse{
		Status: "password reset successful",
	}, nil
}

func (s *Auth) RefreshToken(ctx context.Context, req *authV1.RefreshTokenRequest) (*authV1.RefreshTokenResponse, error) {
	params := &bo.RefreshTokenRequest{
		RefreshToken: req.GetRefreshToken(),
	}
	signIn, err := s.authBiz.RefreshToken(ctx, params)
	if err != nil {
		return nil, err
	}
	return &authV1.RefreshTokenResponse{
		Token:     signIn.Token,
		ExpiresIn: signIn.ExpiresIn,
	}, nil
}

func (s *Auth) ValidateToken(ctx context.Context, token string) (jwt.Claims, error) {
	return s.authBiz.ValidateToken(ctx, token)
}

func (s *Auth) SendSmsCode(ctx context.Context, req *authV1.SendSmsCodeRequest) (*authV1.SendSmsCodeResponse, error) {
	params := &bo.SendSmsConfigParams{
		PhoneNumber: req.GetPhoneNumber(),
	}
	smsCode, err := s.authBiz.SendSmsCode(ctx, params)
	if err != nil {
		return nil, err
	}
	return &authV1.SendSmsCodeResponse{
		ExpireTime: smsCode.ExpireTime,
		Interval:   smsCode.Interval,
	}, nil
}

func (s *Auth) VerifySmsCode(ctx context.Context, req *authV1.VerifySmsCodeRequest) (*authV1.VerifySmsCodeResponse, error) {
	params := &bo.VerifySmsCodeRequest{
		PhoneNumber: req.GetPhoneNumber(),
		VerifyCode:  req.GetVerifyCode(),
	}
	singIn, err := s.authBiz.VerifySmsCode(ctx, params)
	if err != nil {
		return nil, err
	}
	return &authV1.VerifySmsCodeResponse{
		Token:        singIn.Token,
		RefreshToken: singIn.RefreshToken,
		ExpiresIn:    singIn.ExpiresIn,
	}, nil
}

func (s *Auth) OneClickLogin(ctx context.Context, req *authV1.OneClickLoginRequest) (*authV1.OneClickLoginResponse, error) {
	params := &bo.GetMobileConfigParams{
		Token: req.GetToken(),
	}
	signIn, err := s.authBiz.OneClickLogin(ctx, params)
	if err != nil {
		return nil, err
	}
	return &authV1.OneClickLoginResponse{
		Token:        signIn.Token,
		RefreshToken: signIn.RefreshToken,
		ExpiresIn:    signIn.ExpiresIn,
	}, nil
}
