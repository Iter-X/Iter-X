package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	
	"github.com/iter-x/iter-x/internal/common/cnst"
	"github.com/iter-x/iter-x/internal/common/xerr"
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

func ValidToken(secret []byte, s string, claims jwt.Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(s, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})
}

func ExtractClaims(ctx context.Context) (*Claims, error) {
	v := ctx.Value(cnst.CtxKeyClaims)
	claims, ok := v.(*Claims)
	if !ok {
		return nil, xerr.ErrorUnauthorized()
	}
	return claims, nil
}
