package userbusiness

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"
)

// giao dien danh cho storage thuc hien xoa
type GetCardStore interface {
	GetCart(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

// Cau truc cua doi tuong xoa nha hang
type getCardBiz struct {
	store     GetCardStore
	requester common.Requester
}

// Khoi tai biz delete
func NewGetCardBiz(store GetCardStore, requester common.Requester) *getCardBiz {
	return &getCardBiz{store: store, requester: requester}
}

func (biz *getCardBiz) GetCard(context context.Context, id int) (*usermodel.User, error) {
	// oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	user, err := biz.store.GetCart(context, map[string]interface{}{"id": id}, "ListCard")

	if err != nil {
		return nil, common.ErrEntityNotFound(usermodel.EntityName, err)
	}

	if user.Status == 0 {
		return nil, common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	return user, nil
}
