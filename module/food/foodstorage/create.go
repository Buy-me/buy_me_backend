package foodstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/foodmodel"
	"log"
)

func (s *sqlStore) Create(context context.Context, data *foodmodel.FoodCreate) error {

	log.Println("Come here", data)

	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
