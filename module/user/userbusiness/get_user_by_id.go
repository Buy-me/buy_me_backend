package userbusiness

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"
)

// giao dien danh cho storage thuc hien xoa
type GetUserStore interface {
	FindUser(
		context context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
}

// Cau truc cua doi tuong xoa nha hang
type getUserBiz struct {
	store     GetUserStore
	requester common.Requester
}

// Khoi tai biz delete
func NewGetUserBiz(store GetUserStore, requester common.Requester) *getUserBiz {
	return &getUserBiz{store: store, requester: requester}
}

func (biz *getUserBiz) GetUser(context context.Context, id int) (*usermodel.User, error) {
	// oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	user, err := biz.store.FindUser(context, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrEntityNotFound(usermodel.EntityName, err)
	}

	if user.Status == 0 {
		return nil, common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	return user, nil
}
