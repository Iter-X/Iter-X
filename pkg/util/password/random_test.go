package password_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iter-x/iter-x/pkg/util/password"
)

// TestGenerateRandomPassword_Success 验证生成的密码长度是否正确
func TestGenerateRandomPassword_Success(t *testing.T) {
	length := 10
	p := password.GenerateRandomPassword(length)
	assert.Len(t, p, length)
}

// TestGenerateRandomPassword_Randomness 验证生成的密码是否随机
func TestGenerateRandomPassword_Randomness(t *testing.T) {
	length := 10
	password1 := password.GenerateRandomPassword(length)
	password2 := password.GenerateRandomPassword(length)
	assert.NotEqual(t, password1, password2)
}

// TestGenerateRandomPassword_ZeroLength 验证处理 length = 0 的情况
func TestGenerateRandomPassword_ZeroLength(t *testing.T) {
	length := 0
	p := password.GenerateRandomPassword(length)
	assert.Len(t, p, length)
}

// TestGenerateRandomPassword_NegativeLength 验证处理 length < 0 的情况
func TestGenerateRandomPassword_NegativeLength(t *testing.T) {
	length := -1
	p := password.GenerateRandomPassword(length)
	assert.Len(t, p, 0)
}
