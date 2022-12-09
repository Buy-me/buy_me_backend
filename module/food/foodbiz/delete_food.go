package foodbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/foodmodel"
)

// giao dien danh cho storage thuc hien xoa
type DeleteFoodStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*foodmodel.Food, error)
	Delete(context context.Context, id int) error
}

// Cau truc cua doi tuong xoa nha hang
type deleteFoodBiz struct {
	store     DeleteFoodStore
	requester common.Requester
}

// Khoi tai biz delete
func NewDeleteFoodBiz(store DeleteFoodStore, requester common.Requester) *deleteFoodBiz {
	return &deleteFoodBiz{store: store, requester: requester}
}

// Ham Thuc Thi Xoa Nha Hang
func (biz *deleteFoodBiz) DeleteFood(context context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(foodmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(foodmodel.EntityName, nil)
	}

	// if oldData.UserId != biz.requester.GetUserId() {
	// 	return common.ErrNoPermission(nil)
	// }

	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(foodmodel.EntityName, err)
	}

	return nil
}
