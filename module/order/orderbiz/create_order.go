package orderbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/order/ordermodel"
)

type CreateOrderStore interface {
	Create(context context.Context, data *ordermodel.OrderCreate) error
}

type createOrderBiz struct {
	store CreateOrderStore
}

func NewCreateTicketBiz(store CreateOrderStore) *createOrderBiz {
	return &createOrderBiz{store: store}
}

func (biz *createOrderBiz) CreateOrder(context context.Context, data *ordermodel.OrderCreate) error {

	if err := biz.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(ordermodel.EntityName, err)
	}

	return nil
}
