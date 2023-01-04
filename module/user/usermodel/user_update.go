package usermodel

import "food_delivery/common"

type UserUpdate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	Password        string        `json:"-" gorm:"column:password;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Gender          string        `json:"gender" gorm:"column:gender;"`
	BirthDate       string        `json:"birth_date" gorm:"column:birth_date;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Avatar          *common.Image `json:"avatar" gorm:"column:avatar;type:json"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}
