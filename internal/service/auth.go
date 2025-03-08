package service

import (
	"context"

	"github.com/golang-jwt/jwt/v5"

	authV1 "github.com/iter-x/iter-x/internal/api/auth/v1"
	"github.com/iter-x/iter-x/internal/biz"
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
	return s.authBiz.SignIn(ctx, req)
}

func (s *Auth) SignInWithOAuth(ctx context.Context, req *authV1.SignInWithOAuthRequest) (*authV1.SignInWithOAuthResponse, error) {
	token, err := s.authBiz.SignInWithOAuth(ctx, req)
	if err != nil {
		return nil, err
	}
	return &authV1.SignInWithOAuthResponse{
		Token: token,
	}, nil
}

func (s *Auth) SignUp(ctx context.Context, req *authV1.SignUpRequest) (*authV1.SignUpResponse, error) {
	return s.authBiz.SignUp(ctx, req)
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
	return s.authBiz.RefreshToken(ctx, req)
}

func (s *Auth) ValidateToken(ctx context.Context, token string) (jwt.Claims, error) {
	return s.authBiz.ValidateToken(ctx, token)
}

func (s *Auth) SendSmsCode(ctx context.Context, req *authV1.SendSmsCodeRequest) (*authV1.SendSmsCodeResponse, error) {
	return s.authBiz.SendSmsCode(ctx, req)
}

func (s *Auth) VerifySmsCode(ctx context.Context, req *authV1.VerifySmsCodeRequest) (*authV1.VerifySmsCodeResponse, error) {
	return s.authBiz.VerifySmsCode(ctx, req)
}

func (s *Auth) OneClickLogin(ctx context.Context, req *authV1.OneClickLoginRequest) (*authV1.OneClickLoginResponse, error) {
	return s.authBiz.OneClickLogin(ctx, req)
}
