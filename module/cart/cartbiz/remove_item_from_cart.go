package cartbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/cart/cartmodel"
)

// giao dien danh cho storage thuc hien xoa
type DeleteCartStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*cartmodel.Cart, error)
	Delete(context context.Context, userId int, foodId int) error
}

// Cau truc cua doi tuong xoa nha hang
type deleteCartBiz struct {
	store     DeleteCartStore
	requester common.Requester
}

// Khoi tai biz delete
func NewDeleteCartBiz(store DeleteCartStore, requester common.Requester) *deleteCartBiz {
	return &deleteCartBiz{store: store, requester: requester}
}

// Ham Thuc Thi Xoa Nha Hang
func (biz *deleteCartBiz) DeleteCart(context context.Context, userId int, foodId int) error {
	oldData, err := biz.store.
		FindDataWithCondition(context, map[string]interface{}{"food_id": foodId, "user_id": userId})

	if err != nil {
		return common.ErrEntityNotFound(cartmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(cartmodel.EntityName, nil)
	}

	// if oldData.UserId != biz.requester.GetUserId() {
	// 	return common.ErrNoPermission(nil)
	// }

	if err := biz.store.Delete(context, userId, foodId); err != nil {
		return common.ErrCannotDeleteEntity(cartmodel.EntityName, err)
	}

	return nil
}
