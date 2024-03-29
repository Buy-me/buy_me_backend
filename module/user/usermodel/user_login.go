package usermodel

import (
	"errors"
	"food_delivery/common"
)

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrPasswordInvalid = common.NewCustomError(
		errors.New("password invalid"),
		"password invalid",
		"ErrPasswordInvalid",
	)
)
