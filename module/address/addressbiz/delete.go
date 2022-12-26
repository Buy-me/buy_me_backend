package addressbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/address/addressmodel"
	"food_delivery/module/cart/cartmodel"
)

// giao dien danh cho storage thuc hien xoa
type DeleteAddressStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*addressmodel.Address, error)
	Delete(context context.Context, id int) error
}

// Cau truc cua doi tuong xoa nha hang
type deleteAddressBiz struct {
	store     DeleteAddressStore
	requester common.Requester
}

// Khoi tai biz delete
func NewDeleteAddressBiz(store DeleteAddressStore, requester common.Requester) *deleteAddressBiz {
	return &deleteAddressBiz{store: store, requester: requester}
}

// Ham Thuc Thi Xoa Nha Hang
func (biz *deleteAddressBiz) DeleteCart(context context.Context, id int) error {
	oldData, err := biz.store.
		FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(addressmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(addressmodel.EntityName, nil)
	}

	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(cartmodel.EntityName, err)
	}

	return nil
}
