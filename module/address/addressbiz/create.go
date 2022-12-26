package addressbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/address/addressmodel"
	"food_delivery/module/cart/cartmodel"
)

type CreateAddressStore interface {
	Create(context context.Context, data *addressmodel.AddressCreate) error
}

type createAddressBiz struct {
	store CreateAddressStore
}

func NewCreateAddressBiz(store CreateAddressStore) *createAddressBiz {
	return &createAddressBiz{store: store}
}

func (biz *createAddressBiz) CreateAddress(context context.Context, data *addressmodel.AddressCreate) error {

	if err := biz.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(cartmodel.EntityName, err)
	}

	return nil
}
