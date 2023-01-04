package orderbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/order/ordermodel"
)

// giao dien danh cho storage thuc hien xoa
type GetOrderStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*ordermodel.Order, error)
}

// Cau truc cua doi tuong xoa nha hang
type getOrderBiz struct {
	store GetOrderStore
}

// Khoi tai biz delete
func NewGetOrderBiz(store GetOrderStore) *getOrderBiz {
	return &getOrderBiz{store: store}
}

func (biz *getOrderBiz) GetOrder(context context.Context, id int) (*ordermodel.Order, error) {
	// oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	order, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id}, "Items")

	if err != nil {
		return nil, common.ErrEntityNotFound(ordermodel.EntityName, err)
	}

	if order.Status == 0 {
		return nil, common.ErrEntityDeleted(ordermodel.EntityName, nil)
	}

	return order, nil
}
