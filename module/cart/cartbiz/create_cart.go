package cartbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/cart/cartmodel"
)

type CreateCartStore interface {
	Create(context context.Context, data *cartmodel.CartCreate) error
}

type createCartBiz struct {
	store CreateCartStore
}

func NewCreateCartBiz(store CreateCartStore) *createCartBiz {
	return &createCartBiz{store: store}
}

func (biz *createCartBiz) CreateCart(context context.Context, data *cartmodel.CartCreate) error {

	if err := biz.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(cartmodel.EntityName, err)
	}

	return nil
}
