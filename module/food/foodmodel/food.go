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
	Name            string        `json:"name" gorm:"column:name;"`
	Images          *common.Image `json:"images" gorm:"column:images;"`
	Price           float64       `json:"price" gorm:"column:price;"`
	Description     string        `json:"description" gorm:"column:description;"`
}

// type Restaurant struct {
// 	common.SQLModel `json:",inline"`
// 	Name            string             `json:"name" gorm:"column:name;"`
// 	Addr            string             `json:"addr" gorm:"column:addr;"`
// 	Type            RestaurantType     `json:"type" gorm:"column:type;"`
// 	Logo            *common.Image      `json:"logo" gorm:"column:logo;"`
// 	Cover           *common.Images     `json:"cover" gorm:"column:cover;"`
// 	UserId          int                `json:"-" gorm:"column:user_id;"`
// 	User            *common.SimpleUser `json:"user" gorm:"preload:false;"`
// 	LikedCount      int                `json:"liked_count" gorm:"column:liked_count;"`
// }

func (Food) TableName() string {
	return "foods"
}

type FoodCreate struct {
	common.SQLModel `json:",inline"`
	RestaurantId    int           `json:"restaurant_id" gorm:"column:restaurant_id;"`
	Name            string        `json:"name" gorm:"column:name;"`
	Images          *common.Image `json:"images" gorm:"column:images;"`
	Price           float64       `json:"price" gorm:"column:price;"`
	Description     string        `json:"description" gorm:"column:description;"`
}

func (FoodCreate) TableName() string {
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

func (data *FoodCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

// type RestaurantCreate struct {
// 	common.SQLModel `json:",inline"`
// 	Name            string         `json:"name" gorm:"column:name;"`
// 	UserId          int            `json:"-" gorm:"column:user_id;"`
// 	Addr            string         `json:"addr" gorm:"column:addr;"`
// 	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
// 	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
// }

// func (RestaurantCreate) TableName() string {
// 	return Restaurant{}.TableName()
// }

// type RestaurantUpdate struct {
// 	Name  *string        `json:"name" gorm:"column:name;"`
// 	Addr  *string        `json:"addr" gorm:"column:addr;"`
// 	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
// 	Cover *common.Images `json:"cover" gorm:"column:cover;"`
// }

// func (RestaurantUpdate) TableName() string {
// 	return Restaurant{}.TableName()
// }

// // Validate
// func (data *RestaurantCreate) Validate() error {
// 	data.Name = strings.TrimSpace(data.Name)

// 	if data.Name == "" {
// 		return ErrNameIsEmpty
// 	}

// 	return nil

// }

// Mask
// func (r *Food) Mask(isAdminOrOwner bool) {
// 	r.GenUID(common.DbTypeRestaurant)

// 	if u := r.; u != nil {
// 		u.Mask(isAdminOrOwner)
// 	}
// }

// func (data *FoodCreate) Mask(isAdminOrOwner bool) {
// 	data.GenUID(common.DbTypeRestaurant)
// }

// Error
var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
