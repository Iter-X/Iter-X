package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/common/cnst"
	"github.com/iter-x/iter-x/internal/common/xerr"
	"time"
)

type (
	Claims struct {
		jwt.RegisteredClaims

		UID       uuid.UUID `json:"uid"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		Status    bool      `json:"status"`
		AvatarURL string    `json:"avatar_url"`
	}

	AgentClaims struct {
		jwt.RegisteredClaims

		UID uuid.UUID `json:"uid"`
	}
)

// GenerateToken generates a JWT token.
func GenerateToken(secret []byte, claims Claims,
	issuer string, exp time.Duration) (string, error) {
	if len(secret) == 0 {
		return "", jwt.ErrInvalidKey
	}

	now := time.Now()

	claims.Issuer = issuer
	claims.Subject = claims.UID.String()
	claims.ExpiresAt = jwt.NewNumericDate(now.Add(exp))
	claims.NotBefore = jwt.NewNumericDate(now)
	claims.IssuedAt = jwt.NewNumericDate(now)
	claims.ID = uuid.New().String()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func GenerateAgentToken(secret []byte, claims AgentClaims,
	issuer string, exp time.Duration) (string, error) {
	if len(secret) == 0 {
		return "", jwt.ErrInvalidKey
	}

	now := time.Now()

	claims.Issuer = issuer
	claims.Subject = claims.UID.String()
	claims.ExpiresAt = jwt.NewNumericDate(now.Add(exp))
	claims.NotBefore = jwt.NewNumericDate(now)
	claims.IssuedAt = jwt.NewNumericDate(now)
	claims.ID = uuid.New().String()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func ValidToken(secret []byte, s string, claims jwt.Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(s, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})
}

func ExtractClaims[T any](ctx context.Context) (*T, error) {
	v := ctx.Value(cnst.CtxKeyClaims)
	claims, ok := v.(*T)
	if !ok {
		return nil, xerr.ErrorUnauthorized()
	}
	return claims, nil
}
