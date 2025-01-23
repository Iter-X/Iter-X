package auth

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// successfully hashes a valid password
func TestHashPasswordWithValidPassword(t *testing.T) {
	// Arrange
	password := "validPassword123"

	// Act
	hashedPassword, err := HashPassword(password)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

// handles empty password string
func TestHashPasswordWithEmptyPassword(t *testing.T) {
	// Arrange
	password := ""

	// Act
	hashedPassword, err := HashPassword(password)

	// Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrPasswordLength)
	assert.Empty(t, hashedPassword)
}

// handles password string that is too long
func TestHashPasswordWithTooLongPassword(t *testing.T) {
	// Arrange
	longPassword := "a" + strings.Repeat("b", 1000) + "c"

	// Act
	hashedPassword, err := HashPassword(longPassword)

	// Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrPasswordLength)
	assert.Empty(t, hashedPassword)
}

// correct password matches the hash
func TestCorrectPasswordMatchesHash(t *testing.T) {
	// Arrange
	password := "correct_password"
	hash, err := HashPassword(password)
	assert.NoError(t, err)

	// Act
	result := CompareHashAndPassword(password, hash)

	// Assert
	assert.True(t, result)
}

// incorrect password does not match the hash
func TestIncorrectPasswordDoesNotMatchHash(t *testing.T) {
	// Arrange
	password := "correct_password"
	hash, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	// Act
	result := CompareHashAndPassword("incorrect_password", hash)

	// Assert
	assert.False(t, result)
}
