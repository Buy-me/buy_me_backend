package orderbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/order/ordermodel"
)

// giao dien danh cho storage thuc hien xoa
type DeleteOrderStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*ordermodel.Order, error)
	Delete(context context.Context, id int) error
}

// Cau truc cua doi tuong xoa nha hang
type deleteOrderBiz struct {
	store     DeleteOrderStore
	requester common.Requester
}

// Khoi tai biz delete
func NewDeleteOrderBiz(store DeleteOrderStore, requester common.Requester) *deleteOrderBiz {
	return &deleteOrderBiz{store: store, requester: requester}
}

// Ham Thuc Thi Xoa
func (biz *deleteOrderBiz) DeleteOrder(context context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(ordermodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(ordermodel.EntityName, nil)
	}

	// if oldData.UserId != biz.requester.GetUserId() {
	// 	return common.ErrNoPermission(nil)
	// }

	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(ordermodel.EntityName, err)
	}

	return nil
}
