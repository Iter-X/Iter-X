package password_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iter-x/iter-x/pkg/util/password"
)

// TestNew_DefaultSalt 验证使用默认盐值创建密码对象
func TestNew_DefaultSalt(t *testing.T) {
	pwd := "mySecretPassword"
	passwordObj := password.New(pwd)
	assert.NotNil(t, passwordObj)
	assert.NotEmpty(t, passwordObj.Salt())
	assert.Equal(t, pwd, passwordObj.PValue())
}

// TestNew_CustomSalt 验证使用自定义盐值创建密码对象
func TestNew_CustomSalt(t *testing.T) {
	pwd := "mySecretPassword"
	customSalt := "customSaltValue"
	passwordObj := password.New(pwd, customSalt)
	assert.NotNil(t, passwordObj)
	assert.Equal(t, customSalt, passwordObj.Salt())
	assert.Equal(t, pwd, passwordObj.PValue())
}

// TestEQ_Success 验证正确的哈希密码是否匹配
func TestEQ_Success(t *testing.T) {
	pwd := "mySecretPassword"
	passwordObj := password.New(pwd)
	hashedPassword, _ := passwordObj.EnValue()
	assert.True(t, passwordObj.EQ(hashedPassword))
}

// TestEQ_Failure 验证错误的哈希密码是否不匹配
func TestEQ_Failure(t *testing.T) {
	pwd := "mySecretPassword"
	passwordObj := password.New(pwd)
	assert.False(t, passwordObj.EQ("wrongHashedPassword"))
}

// TestEQ_EmptyHashedPassword 验证空哈希密码是否不匹配
func TestEQ_EmptyHashedPassword(t *testing.T) {
	pwd := "mySecretPassword"
	passwordObj := password.New(pwd)
	assert.False(t, passwordObj.EQ(""))
}

// TestPValue_Success 验证返回的原始密码是否正确
func TestPValue_Success(t *testing.T) {
	pwd := "mySecretPassword"
	passwordObj := password.New(pwd)
	assert.Equal(t, pwd, passwordObj.PValue())
}

// TestEnValue_Success 验证加密后的密码是否正确
func TestEnValue_Success(t *testing.T) {
	pwd := "mySecretPassword"
	passwordObj := password.New(pwd)
	encryptedPassword, err := passwordObj.EnValue()
	assert.NoError(t, err)
	assert.NotEmpty(t, encryptedPassword)
}

// TestSalt_Success 验证返回的盐值是否正确
func TestSalt_Success(t *testing.T) {
	pwd := "mySecretPassword"
	passwordObj := password.New(pwd)
	assert.NotEmpty(t, passwordObj.Salt())
}
