package foodmodel

import (
	"food_delivery/common"
)

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}
type FoodRating struct {
	common.SQLModel `json:",inline"`
	UserId          int     `json:"user_id" gorm:"column:user_id;"`
	User            User    `json:"user" gorm:"-"`
	FoodId          int     `json:"food_id" gorm:"column:food_id;"`
	Rating          float64 `json:"rating" gorm:"column:point;"`
	Comment         string  `json:"comment" gorm:"column:comment;"`
}

func (FoodRating) TableName() string {
	return "food_ratings"
}

type FoodRatingCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int     `json:"-" gorm:"column:user_id;"`
	FoodId          int     `json:"-" gorm:"column:food_id;"`
	Rating          float64 `json:"rating" gorm:"column:point;"`
	Comment         string  `json:"comment" gorm:"column:comment;"`
}

func (FoodRatingCreate) TableName() string {
	return "food_ratings"
}
