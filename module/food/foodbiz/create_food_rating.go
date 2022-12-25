package foodbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/foodmodel"
)

type CreateFoodRatingStore interface {
	CreateFoodRating(context context.Context, data *foodmodel.FoodRating) error
}

type createFoodRatingBiz struct {
	store CreateFoodRatingStore
}

func NewCreateFoodRatingBiz(store CreateFoodRatingStore) *createFoodRatingBiz {
	return &createFoodRatingBiz{store: store}
}

func (biz *createFoodRatingBiz) CreateFoodRating(context context.Context, data *foodmodel.FoodRating) error {

	if err := biz.store.CreateFoodRating(context, data); err != nil {
		return common.ErrCannotCreateEntity(foodmodel.EntityName, err)
	}
	return nil
}
