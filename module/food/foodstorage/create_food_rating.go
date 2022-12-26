package foodstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/foodmodel"
	"log"
)

func (s *sqlStore) CreateFoodRating(context context.Context, data *foodmodel.FoodRating) error {

	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	go func(data *foodmodel.FoodRating) {
		var food foodmodel.Food

		if err := s.db.Where("id = ?", data.FoodId).First(&food).Error; err != nil {
			log.Println(err)
			return
		}

		log.Println(food.Rating, food.CountRating)

		newCountRating := food.CountRating + 1
		newRating := ((food.Rating * float64(food.CountRating)) + data.Rating) / (float64(newCountRating))

		if err := s.db.Table(foodmodel.Food{}.TableName()).Where("id = ?", data.FoodId).Updates(map[string]interface{}{"rating": newRating, "count_rating": newCountRating}).Error; err != nil {
			return
		}

	}(data)

	return nil
}
