package userbusiness

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"
)

// giao dien danh cho storage thuc hien xoa
type GetFavouriteStore interface {
	GetFavourite(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

// Cau truc cua doi tuong xoa nha hang
type getFavouriteBiz struct {
	store     GetFavouriteStore
	requester common.Requester
}

// Khoi tai biz delete
func NewGetFavouriteBiz(store GetFavouriteStore, requester common.Requester) *getFavouriteBiz {
	return &getFavouriteBiz{store: store, requester: requester}
}

func (biz *getFavouriteBiz) GetFavourite(context context.Context, id int) (*usermodel.User, error) {
	// oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	user, err := biz.store.GetFavourite(context, map[string]interface{}{"id": id}, "ListFavourite")

	if err != nil {
		return nil, common.ErrEntityNotFound(usermodel.EntityName, err)
	}

	if user.Status == 0 {
		return nil, common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	return user, nil
}
