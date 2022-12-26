package addressmodel

import (
	"errors"
	"food_delivery/common"
)

const EntityName = "Address"

type Address struct {
	common.SQLModel `json:",inline"`
	UserId          int    `json:"-" gorm:"column:user_id;"`
	Title           string `json:"title" gorm:"column:title;"`
	Address         string `json:"address" gorm:"column:addr;"`
}

func (Address) TableName() string {
	return "user_addresses"
}

type AddressCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int    `json:"-" gorm:"column:user_id;"`
	Title           string `json:"title" gorm:"column:title;"`
	Address         string `json:"address" gorm:"column:addr;"`
}

func (AddressCreate) TableName() string {
	return Address{}.TableName()
}

// Error
var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
