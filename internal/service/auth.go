package service

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iter-x/iter-x/internal/api/auth/v1"
	"github.com/iter-x/iter-x/internal/biz"
)

type Auth struct {
	v1.UnimplementedAuthServiceServer
	authBiz *biz.Auth
}

func NewAuth(authBiz *biz.Auth) *Auth {
	return &Auth{
		authBiz: authBiz,
	}
}

func (s *Auth) SignIn(ctx context.Context, req *v1.SignInRequest) (*v1.SignInResponse, error) {
	return s.authBiz.SignIn(ctx, req)
}

func (s *Auth) SignInWithOAuth(ctx context.Context, req *v1.SignInWithOAuthRequest) (*v1.SignInWithOAuthResponse, error) {
	token, err := s.authBiz.SignInWithOAuth(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.SignInWithOAuthResponse{
		Token: token,
	}, nil
}

func (s *Auth) SignUp(ctx context.Context, req *v1.SignUpRequest) (*v1.SignUpResponse, error) {
	return s.authBiz.SignUp(ctx, req)
}

func (s *Auth) RequestPasswordReset(ctx context.Context, req *v1.RequestPasswordResetRequest) (*v1.RequestPasswordResetResponse, error) {
	// TODO: Implement logic for requesting a password reset
	return &v1.RequestPasswordResetResponse{
		Status: "password reset request sent",
	}, nil
}

func (s *Auth) VerifyPasswordResetToken(ctx context.Context, req *v1.VerifyPasswordResetTokenRequest) (*v1.VerifyPasswordResetTokenResponse, error) {
	// TODO: Implement logic for verifying the password reset token
	return &v1.VerifyPasswordResetTokenResponse{
		Valid: true,
	}, nil
}

func (s *Auth) ResetPassword(ctx context.Context, req *v1.ResetPasswordRequest) (*v1.ResetPasswordResponse, error) {
	// TODO: Implement logic for resetting the password
	return &v1.ResetPasswordResponse{
		Status: "password reset successful",
	}, nil
}

func (s *Auth) RefreshToken(ctx context.Context, req *v1.RefreshTokenRequest) (*v1.RefreshTokenResponse, error) {
	return s.authBiz.RefreshToken(ctx, req)
}

func (s *Auth) ValidateToken(ctx context.Context, token string) (jwt.Claims, error) {
	return s.authBiz.ValidateToken(ctx, token)
}

// GetSmsAuthTokens get sms auth tokens
func (s *Auth) GetSmsAuthTokens(ctx context.Context, req *v1.GetSmsAuthTokensRequest) (*v1.GetSmsAuthTokensResponse, error) {
	return s.authBiz.GetSmsAuthTokens(ctx, req)
}

// VerifySmsCode verify sms code
func (s *Auth) VerifySmsCode(ctx context.Context, req *v1.VerifySmsCodeRequest) (*v1.VerifySmsCodeResponse, error) {
	return s.authBiz.VerifySmsCode(ctx, req)
}
