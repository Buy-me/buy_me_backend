package foodbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/foodmodel"
)

// giao dien danh cho storage thuc hien xoa
type UpdateFoodStore interface {
	Update(ctx context.Context, data foodmodel.FoodUpdate, id int) error
}

// Cau truc cua doi tuong xoa nha hang
type updateFoodBiz struct {
	store     UpdateFoodStore
	requester common.Requester
}

// Khoi tai biz delete
func NewUpdateFoodBiz(store UpdateFoodStore, requester common.Requester) *updateFoodBiz {
	return &updateFoodBiz{store: store, requester: requester}
}

func (biz *updateFoodBiz) UpdateFood(context context.Context, data foodmodel.FoodUpdate, id int) error {
	// oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	err := biz.store.Update(context, data, id)

	if err != nil {
		return err
	}

	return nil
}
