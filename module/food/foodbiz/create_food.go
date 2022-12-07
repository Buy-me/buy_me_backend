package foodbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/foodmodel"
)

type CreateFoodStore interface {
	Create(context context.Context, data *foodmodel.FoodCreate) error
}

type createFoodBiz struct {
	store CreateFoodStore
}

func NewCreateFoodBiz(store CreateFoodStore) *createFoodBiz {
	return &createFoodBiz{store: store}
}

func (biz *createFoodBiz) CreateFood(context context.Context, data *foodmodel.FoodCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := biz.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(foodmodel.EntityName, err)
	}

	return nil
}
