package userbusiness

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"
)

// giao dien danh cho storage thuc hien xoa
type GetAddressStore interface {
	GetCart(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

// Cau truc cua doi tuong xoa nha hang
type getAddressBiz struct {
	store     GetAddressStore
	requester common.Requester
}

// Khoi tai biz delete
func NewGetAddressBiz(store GetAddressStore, requester common.Requester) *getAddressBiz {
	return &getAddressBiz{store: store, requester: requester}
}

func (biz *getAddressBiz) GetAddresses(context context.Context, id int) (*usermodel.User, error) {
	// oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	user, err := biz.store.GetCart(context, map[string]interface{}{"id": id}, "ListAddress")

	if err != nil {
		return nil, common.ErrEntityNotFound(usermodel.EntityName, err)
	}

	if user.Status == 0 {
		return nil, common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	return user, nil
}
