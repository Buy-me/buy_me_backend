package cardmodel

import (
	"errors"
	"food_delivery/common"
)

const EntityName = "Card"

type Card struct {
	common.SQLModel `json:",inline"`
	UserId          int    `json:"-" gorm:"column:user_id;"`
	ExpireDate      string `json:"expire_date" gorm:"column:expire_date;"`
	CVV             string `json:"cvv" gorm:"column:cvv;"`
	Number          string `json:"number" gorm:"column:number;"`
	Name            string `json:"name" gorm:"column:name;"`
	TypeCard        string `json:"type_card" gorm:"column:type_card;"`
}

func (Card) TableName() string {
	return "user_cards"
}

type CardCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int    `json:"-" gorm:"column:user_id;"`
	ExpireDate      string `json:"expire_date" gorm:"column:expire_date;"`
	CVV             string `json:"cvv" gorm:"column:cvv;"`
	Number          string `json:"number" gorm:"column:number;"`
	Name            string `json:"name" gorm:"column:name;"`
	TypeCard        string `json:"type_card" gorm:"column:type_card;"`
}

func (CardCreate) TableName() string {
	return Card{}.TableName()
}

// Error
var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
