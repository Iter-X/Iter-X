package bo

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/pkg/util/password"
	"github.com/iter-x/iter-x/pkg/vobj"
)

type (
	SignInResponse struct {
		Token        string
		RefreshToken string
		ExpiresIn    float64
	}

	CreateUserByPhoneParam struct {
		PhoneNumber string
	}
)

// GenerateUserDo generate user do
func (c *CreateUserByPhoneParam) GenerateUserDo() *do.User {
	passwordObj := password.New(password.GenerateRandomPassword(8))
	enPass, err := passwordObj.EnValue()
	if err != nil {
		panic(err)
	}
	return &do.User{
		Status:   vobj.UserStatusActive,
		Username: c.PhoneNumber,
		Password: enPass,
		Salt:     passwordObj.Salt(),
		Phone:    c.PhoneNumber,
	}
}
