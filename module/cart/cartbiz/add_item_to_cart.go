package cartbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/cart/cartmodel"
)

type CreateCartStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*cartmodel.Cart, error)
	Create(context context.Context, data *cartmodel.CartCreate) error
}

type createCartBiz struct {
	store CreateCartStore
}

func NewCreateCartBiz(store CreateCartStore) *createCartBiz {
	return &createCartBiz{store: store}
}

func (biz *createCartBiz) CreateCart(context context.Context, data *cartmodel.CartCreate) error {
	oldData, _ := biz.store.
		FindDataWithCondition(context, map[string]interface{}{"food_id": data.FoodId, "user_id": data.UserId})

	if oldData != nil {
		return common.ErrEntityExisted(cartmodel.EntityName+" items", nil)
	}

	if err := biz.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(cartmodel.EntityName, err)
	}

	return nil
}
