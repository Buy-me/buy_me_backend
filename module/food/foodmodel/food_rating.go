package foodmodel

import (
	"food_delivery/common"
)

type FoodRating struct {
	common.SQLModel `json:",inline"`
	UserId          int     `json:"user_id" gorm:"column:user_id;"`
	FoodId          int     `json:"food_id" gorm:"column:food_id;"`
	Rating          float64 `json:"rating" gorm:"column:point;"`
	Comment         string  `json:"comment" gorm:"column:comment;"`
}

func (FoodRating) TableName() string {
	return "food_ratings"
}
