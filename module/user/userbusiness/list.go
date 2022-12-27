package userbusiness

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"
)

// Dinh Nghia giao dien cho storage thi hanh
type ListUserStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *usermodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]usermodel.User, error)
}

// Dinh nghia cau truc nghiep vu nha hang
type listUserBiz struct {
	store ListUserStore
}

// Khoi tao doi tuong lay danh sach nha hang
func NewListUserBiz(store ListUserStore) *listUserBiz {
	return &listUserBiz{store: store}
}

// Dinh Nghia Chuc Nang Cho Transport Su Dung
func (biz *listUserBiz) ListUser(
	context context.Context,
	filter *usermodel.Filter,
	paging *common.Paging,
) ([]usermodel.User, error) {

	result, err := biz.store.ListDataWithCondition(context, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
