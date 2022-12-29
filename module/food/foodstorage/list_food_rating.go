package foodstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/foodmodel"
)

func (s *sqlStore) ListFoodRating(context context.Context, foodId int, moreKeys ...string) ([]foodmodel.FoodRating, error) {
	var result []foodmodel.FoodRating

	//or (1,2,3)
	db := s.db.Table(foodmodel.FoodRating{}.TableName()).Where("status in (1)")

	if foodId > 0 {
		db = db.Where("food_id = ?", foodId)
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
