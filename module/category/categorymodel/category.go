package categorymodel

import (
	"errors"
	"food_delivery/common"
)

const EntityName = "Category"

type Category struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Description     string `json:"description" gorm:"column:description;"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Description     string `json:"description" gorm:"column:description;"`
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
