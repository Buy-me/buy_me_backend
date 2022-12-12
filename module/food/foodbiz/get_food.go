package foodbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/foodmodel"
)

// giao dien danh cho storage thuc hien xoa
type GetFoodStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*foodmodel.Food, error)
}

// Cau truc cua doi tuong xoa nha hang
type getFoodBiz struct {
	store     GetFoodStore
	requester common.Requester
}

// Khoi tai biz delete
func NewGetFoodBiz(store GetFoodStore, requester common.Requester) *getFoodBiz {
	return &getFoodBiz{store: store, requester: requester}
}

func (biz *getFoodBiz) GetFood(context context.Context, id int) (*foodmodel.Food, error) {
	// oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	food, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrEntityNotFound(foodmodel.EntityName, err)
	}

	if food.Status == 0 {
		return nil, common.ErrEntityDeleted(foodmodel.EntityName, nil)
	}

	return food, nil
}
