package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

// Generates a valid JWT token with given payload and expiration
func TestGenerateTokenWithPayloadAndExpiration(t *testing.T) {
	// Arrange
	secret := []byte("my-secret")
	issuer := "test-issuer"
	exp := time.Hour
	payload := Claims{Username: "test-user"}

	// Act
	token, err := GenerateToken(secret, payload, issuer, exp)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

// Handles nil payload by initializing an empty claims
func TestGenerateTokenWithNilPayload(t *testing.T) {
	// Arrange
	secret := []byte("my-secret")
	issuer := "test-issuer"
	exp := time.Hour
	var payload Claims

	// Act
	token, err := GenerateToken(secret, payload, issuer, exp)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestGenerateTokenWithEmptySecret(t *testing.T) {
	// Arrange
	var secret []byte
	issuer := "test-issuer"
	exp := time.Hour
	payload := Claims{Username: "test-user"}

	// Act
	token, err := GenerateToken(secret, payload, issuer, exp)

	// Assert
	assert.ErrorIs(t, err, jwt.ErrInvalidKey)
	assert.Empty(t, token)
}
