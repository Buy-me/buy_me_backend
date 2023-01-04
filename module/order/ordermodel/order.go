package ordermodel

import (
	"errors"
	"food_delivery/common"
)

type RestaurantType string

const EntityName = "Order"

type Order struct {
	common.SQLModel `json:",inline"`
	UserId          int           `json:"user_id" gorm:"column:user_id;"`
	Name            string        `json:"name" gorm:"column:name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	TitleAddress    string        `json:"title_address" gorm:"column:title;"`
	DetailAddress   string        `json:"detail_address" gorm:"column:addr;"`
	TotalPrice      float64       `json:"total_price" gorm:"column:total_price;"`
	State           string        `json:"tracking_state" gorm:"column:state;default:pending;"`
	Items           []OrderDetail `json:"items"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int                 `json:"user_id" gorm:"column:user_id;"`
	Name            string              `json:"name" gorm:"column:name;"`
	Phone           string              `json:"phone" gorm:"column:phone;"`
	TitleAddress    string              `json:"title_address" gorm:"column:title;"`
	DetailAddress   string              `json:"detail_address" gorm:"column:addr;"`
	TotalPrice      float64             `json:"total_price" gorm:"column:total_price;"`
	State           string              `json:"tracking_state" gorm:"column:state;default:pending;"`
	Items           []OrderDetailCreate `json:"items"`
}

func (OrderCreate) TableName() string {
	return Order{}.TableName()
}

// Mask
func (r *Order) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeOrder)
}

func (data *OrderCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeOrder)
}

// Error
var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
