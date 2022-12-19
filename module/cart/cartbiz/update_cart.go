package cartbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/cart/cartmodel"
)

// giao dien danh cho storage thuc hien xoa
type UpdateCartStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*cartmodel.Cart, error)
	Update(ctx context.Context, quantity int, userId int, foodId int) error
}

// Cau truc cua doi tuong xoa nha hang
type updateCartBiz struct {
	store     UpdateCartStore
	requester common.Requester
}

// Khoi tai biz delete
func NewUpdateCartBiz(store UpdateCartStore, requester common.Requester) *updateCartBiz {
	return &updateCartBiz{store: store, requester: requester}
}

func (biz *updateCartBiz) UpdateCart(ctx context.Context, quantity int, userId int, foodId int) error {
	// oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	_, err := biz.store.
		FindDataWithCondition(ctx, map[string]interface{}{"food_id": foodId, "user_id": userId})

	if err != nil {
		return common.ErrEntityNotFound(cartmodel.EntityName+" item", err)
	}

	if err := biz.store.Update(ctx, quantity, userId, foodId); err != nil {
		return err
	}

	return nil
}
