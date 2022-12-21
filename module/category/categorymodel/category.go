package categorymodel

import (
	"errors"
	"food_delivery/common"
)

const EntityName = "Category"

type Food struct {
	common.SQLModel `json:",inline"`
	RestaurantId    int           `json:"restaurant_id" gorm:"column:restaurant_id;"`
	CategoryId      int           `json:"category_id" gorm:"column:category_id;"`
	Name            string        `json:"name" gorm:"column:name;"`
	Images          *common.Image `json:"images" gorm:"column:images;"`
	Price           float64       `json:"price" gorm:"column:price;"`
	Description     string        `json:"description" gorm:"column:description;"`
}

type Category struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name;"`
	Description     string        `json:"description" gorm:"column:description;"`
	Icon            *common.Image `json:"icon" gorm:"column:icon;"`
	Food            []Food        `json:"foods" gorm:"-"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryCreate struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name;"`
	Description     string        `json:"description" gorm:"column:description;"`
	Icon            *common.Image `json:"icon" gorm:"column:icon;"`
}

func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}

// Mask
func (r *Category) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeCategory)
}

func (data *CategoryCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeCategory)
}

// Error
var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
