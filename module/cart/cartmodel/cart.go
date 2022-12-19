package cartmodel

import (
	"errors"
	"food_delivery/module/food/foodmodel"
	"time"
)

const EntityName = "Cart"

type Cart struct {
	UserId    int             `json:"-" gorm:"column:user_id;"`
	FoodId    int             `json:"food_id" gorm:"column:food_id;"`
	Food      *foodmodel.Food `json:"food_data" gorm:"-"`
	Quantity  int             `json:"quantity" gorm:"column:quantity;"`
	Status    int             `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time      `json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty" gorm:"updated_at"`
}

func (Cart) TableName() string {
	return "carts"
}

type CartCreate struct {
	UserId    int        `json:"-" gorm:"column:user_id;"`
	FoodId    int        `json:"food_id" gorm:"column:food_id;"`
	Quantity  int        `json:"quantity" gorm:"column:quantity;"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"updated_at"`
}

func (CartCreate) TableName() string {
	return Cart{}.TableName()
}

// Error
var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
