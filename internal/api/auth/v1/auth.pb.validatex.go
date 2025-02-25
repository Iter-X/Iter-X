// Code generated by protoc-gen-validatex. DO NOT EDIT.

// versions:
//  protoc-gen-validatex v0.7.0

package v1

import (
	context "context"
	i18n "github.com/nicksnyder/go-i18n/v2/i18n"
	validatex "github.com/protoc-gen/protoc-gen-validatex/pkg/validatex"
)

func (x *SignInRequest) Validate(ctx context.Context) error {
	if x == nil {
		return nil
	}
	if validatex.ValidEmail(x.Email) != nil {
		return validatex.NewError(
			validatex.MustLocalize(ctx, &i18n.LocalizeConfig{MessageID: "EmailInvalid",
				TemplateData: map[string]string{"FieldName": "email"},
			}, "must be a valid email")).
			WithMetadata(map[string]string{"field": "email"})
	}
	if len(x.Password) < 5 {
		return validatex.NewError(
			validatex.MustLocalize(ctx, &i18n.LocalizeConfig{MessageID: "StringMinLen",
				TemplateData: map[string]string{"MinLen": "5"},
			}, "must be at least 5 characters long")).
			WithMetadata(map[string]string{"field": "password"})
	}
	if len(x.Password) > 50 {
		return validatex.NewError(
			validatex.MustLocalize(ctx, &i18n.LocalizeConfig{MessageID: "StringMaxLen",
				TemplateData: map[string]string{"MaxLen": "50"},
			}, "must be at most 50 characters long")).
			WithMetadata(map[string]string{"field": "password"})
	}
	return nil
}
func (x *SignInResponse) Validate(ctx context.Context) error {
	if x == nil {
		return nil
	}
	return nil
}
func (x *SignInWithOAuthRequest) Validate(ctx context.Context) error {
	if x == nil {
		return nil
	}
	if len(x.Code) < 1 {
		return validatex.NewError(
			validatex.MustLocalize(ctx, &i18n.LocalizeConfig{MessageID: "StringMinLen",
				TemplateData: map[string]string{"MinLen": "1"},
			}, "must be at least 1 characters long")).
			WithMetadata(map[string]string{"field": "code"})
	}
	if len(x.Code) > 256 {
		return validatex.NewError(
			validatex.MustLocalize(ctx, &i18n.LocalizeConfig{MessageID: "StringMaxLen",
				TemplateData: map[string]string{"MaxLen": "256"},
			}, "must be at most 256 characters long")).
			WithMetadata(map[string]string{"field": "code"})
	}
	return nil
}
func (x *SignInWithOAuthResponse) Validate(ctx context.Context) error {
	if x == nil {
		return nil
	}
	return nil
}
func (x *SignUpRequest) Validate(ctx context.Context) error {
	if x == nil {
		return nil
	}
	if validatex.ValidEmail(x.Email) != nil {
		return validatex.NewError(
			validatex.MustLocalize(ctx, &i18n.LocalizeConfig{MessageID: "EmailInvalid",
				TemplateData: map[string]string{"FieldName": "email"},
			}, "must be a valid email")).
			WithMetadata(map[string]string{"field": "email"})
	}
	if len(x.Password) < 5 {
		return validatex.NewError(
			validatex.MustLocalize(ctx, &i18n.LocalizeConfig{MessageID: "StringMinLen",
				TemplateData: map[string]string{"MinLen": "5"},
			}, "must be at least 5 characters long")).
			WithMetadata(map[string]string{"field": "password"})
	}
	if len(x.Password) > 50 {
		return validatex.NewError(
			validatex.MustLocalize(ctx, &i18n.LocalizeConfig{MessageID: "StringMaxLen",
				TemplateData: map[string]string{"MaxLen": "50"},
			}, "must be at most 50 characters long")).
			WithMetadata(map[string]string{"field": "password"})
	}
	return nil
}
func (x *SignUpResponse) Validate(ctx context.Context) error {
	if x == nil {
		return nil
	}
	return nil
}
func (x *RequestPasswordResetRequest) Validate(ctx context.Context) error {
	if x == nil {
		return nil
	}
	return nil
}
func (x *RequestPasswordResetResponse) Validate(ctx context.Context) error {
	if x == nil {
		return nil
	}
	return nil
}
func (x *VerifyPasswordResetTokenRequest) Validate(ctx context.Context) error {
	if x == nil {
		return nil
	}
	return nil
}
func (x *VerifyPasswordResetTokenResponse) Validate(ctx context.Context) error {
	if x == nil {
		return nil
	}
	return nil
}
func (x *ResetPasswordRequest) Validate(ctx context.Context) error {
	if x == nil {
		return nil
	}
	return nil
}
func (x *ResetPasswordResponse) Validate(ctx context.Context) error {
	if x == nil {
		return nil
	}
	return nil
}
func (x *RefreshTokenRequest) Validate(ctx context.Context) error {
	if x == nil {
		return nil
	}
	if len(x.RefreshToken) == 0 {
		return validatex.NewError(
			validatex.MustLocalize(ctx, &i18n.LocalizeConfig{MessageID: "StringNonEmpty"}, "must not be empty")).
			WithMetadata(map[string]string{"field": "refresh_token"})
	}
	return nil
}
func (x *RefreshTokenResponse) Validate(ctx context.Context) error {
	if x == nil {
		return nil
	}
	return nil
}

func init() {
	validatex.Init18n("./i18n/validatex")
}
