package userbusiness

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"
)

// giao dien danh cho storage thuc hien xoa
type UpdateProfileStore interface {
	Update(ctx context.Context, data usermodel.UserUpdate, id int) error
}

// Cau truc cua doi tuong xoa nha hang
type updateProfileBiz struct {
	store     UpdateProfileStore
	requester common.Requester
}

// Khoi tai biz delete
func NewUpdateProfileBiz(store UpdateProfileStore, requester common.Requester) *updateProfileBiz {
	return &updateProfileBiz{store: store, requester: requester}
}

func (biz *updateProfileBiz) UpdateProfile(context context.Context, data usermodel.UserUpdate, id int) error {
	// oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	err := biz.store.Update(context, data, id)

	if err != nil {
		return err
	}

	return nil
}
