package service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	pb "github.com/iter-x/iter-x/internal/api/auth/v1"
	"github.com/iter-x/iter-x/internal/biz"
)

type Auth struct {
	pb.UnimplementedAuthServiceServer
	authBiz *biz.Auth
}

func NewAuth(authBiz *biz.Auth) *Auth {
	return &Auth{
		authBiz: authBiz,
	}
}

func (s *Auth) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	return s.authBiz.SignIn(ctx, req)
}

func (s *Auth) SignInWithOAuth(ctx context.Context, req *pb.SignInWithOAuthRequest) (*pb.SignInWithOAuthResponse, error) {
	token, err := s.authBiz.SignInWithOAuth(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SignInWithOAuthResponse{
		Token: token,
	}, nil
}

func (s *Auth) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return s.authBiz.SignUp(ctx, req)
}

func (s *Auth) RequestPasswordReset(ctx context.Context, req *pb.RequestPasswordResetRequest) (*pb.RequestPasswordResetResponse, error) {
	// TODO: Implement logic for requesting a password reset
	return &pb.RequestPasswordResetResponse{
		Status: "password reset request sent",
	}, nil
}

func (s *Auth) VerifyPasswordResetToken(ctx context.Context, req *pb.VerifyPasswordResetTokenRequest) (*pb.VerifyPasswordResetTokenResponse, error) {
	// TODO: Implement logic for verifying the password reset token
	return &pb.VerifyPasswordResetTokenResponse{
		Valid: true,
	}, nil
}

func (s *Auth) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	// TODO: Implement logic for resetting the password
	return &pb.ResetPasswordResponse{
		Status: "password reset successful",
	}, nil
}

func (s *Auth) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	return s.authBiz.RefreshToken(ctx, req)
}

func (s *Auth) ValidateToken(ctx context.Context, token string) (jwt.Claims, error) {
	return s.authBiz.ValidateToken(ctx, token)
}
