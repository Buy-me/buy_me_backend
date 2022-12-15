package foodmodel

import (
	"errors"
	"food_delivery/common"
	"strings"
)

type RestaurantType string

const EntityName = "Food"

type Food struct {
	common.SQLModel `json:",inline"`
	RestaurantId    int           `json:"restaurant_id" gorm:"column:restaurant_id;"`
	CategoryId      int           `json:"category_id" gorm:"column:category_id;"`
	Name            string        `json:"name" gorm:"column:name;"`
	Images          *common.Image `json:"images" gorm:"column:images;"`
	Price           float64       `json:"price" gorm:"column:price;"`
	Description     string        `json:"description" gorm:"column:description;"`
}

func (Food) TableName() string {
	return "foods"
}

type FoodCreate struct {
	common.SQLModel `json:",inline"`
	RestaurantId    int           `json:"restaurant_id" gorm:"column:restaurant_id;"`
	CategoryId      int           `json:"category_id" gorm:"column:category_id;"`
	Name            string        `json:"name" gorm:"column:name;"`
	Images          *common.Image `json:"images" gorm:"column:images;"`
	Price           float64       `json:"price" gorm:"column:price;"`
	Description     string        `json:"description" gorm:"column:description;"`
}

func (FoodCreate) TableName() string {
	return Food{}.TableName()
}

type FoodUpdate struct {
	Name        string        `json:"name" gorm:"column:name;"`
	Images      *common.Image `json:"images" gorm:"column:images;"`
	Price       float64       `json:"price" gorm:"column:price;"`
	Description string        `json:"description" gorm:"column:description;"`
}

func (FoodUpdate) TableName() string {
	return Food{}.TableName()
}

// Validate
func (data *FoodCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}

// Mask
func (r *Food) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeFood)
}

func (data *FoodCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeFood)
}

// Error
var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
