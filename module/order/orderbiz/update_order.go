package orderbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/order/ordermodel"
)

// giao dien danh cho storage thuc hien xoa
type UpdateOrderStore interface {
	Update(ctx context.Context, data ordermodel.OrderUpdate, id int) error
}

// Cau truc cua doi tuong xoa nha hang
type updateOrderBiz struct {
	store     UpdateOrderStore
	requester common.Requester
}

// Khoi tai biz delete
func NewUpdateOrderBiz(store UpdateOrderStore, requester common.Requester) *updateOrderBiz {
	return &updateOrderBiz{store: store, requester: requester}
}

func (biz *updateOrderBiz) UpdateFood(context context.Context, data ordermodel.OrderUpdate, id int) error {
	// oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	err := biz.store.Update(context, data, id)

	if err != nil {
		return err
	}

	return nil
}
