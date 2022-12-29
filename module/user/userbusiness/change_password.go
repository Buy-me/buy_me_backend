package userbusiness

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"
)

type ChangePasswordStore interface {
	Update(ctx context.Context, data usermodel.UserUpdate, id int) error
}

type changePasswordBiz struct {
	store     ChangePasswordStore
	requester common.Requester
	hasher    Hasher
}

func NewChangePasswordBiz(store ChangePasswordStore, requester common.Requester, hasher Hasher) *changePasswordBiz {
	return &changePasswordBiz{store: store, requester: requester, hasher: hasher}
}

func (biz *changePasswordBiz) ChangePassword(context context.Context, data usermodel.UserUpdate, id int) error {
	// oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	
	err := biz.store.Update(context, data, id)

	if err != nil {
		return err
	}

	return nil
}
