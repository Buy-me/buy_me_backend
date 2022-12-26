package favouritemodel

import (
	"errors"
	"food_delivery/module/food/foodmodel"
	"time"
)

const EntityName = "Favourite"

type Favourite struct {
	UserId    int             `json:"-" gorm:"column:user_id;"`
	FoodId    int             `json:"food_id" gorm:"column:food_id;"`
	Food      *foodmodel.Food `json:"food_data" gorm:"-"`
	Status    int             `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time      `json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty" gorm:"updated_at"`
}

func (Favourite) TableName() string {
	return "favourites"
}

type FavouriteCreate struct {
	UserId    int        `json:"-" gorm:"column:user_id;"`
	FoodId    int        `json:"food_id" gorm:"column:food_id;"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"updated_at"`
}

func (FavouriteCreate) TableName() string {
	return Favourite{}.TableName()
}

// Error
var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
