package userbusiness

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"
)

// giao dien danh cho storage thuc hien xoa
type DeleteUserStore interface {
	Delete(context context.Context, id int) error
}

// Cau truc cua doi tuong xoa nha hang
type deleteUserBiz struct {
	store     DeleteUserStore
	requester common.Requester
}

// Khoi tai biz delete
func NewDeleteUserBiz(store DeleteUserStore, requester common.Requester) *deleteUserBiz {
	return &deleteUserBiz{store: store, requester: requester}
}

// Ham Thuc Thi Xoa Nha Hang
func (biz *deleteUserBiz) DeleteFood(context context.Context, id int) error {

	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(usermodel.EntityName, err)
	}

	return nil
}
