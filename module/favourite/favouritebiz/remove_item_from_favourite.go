package favouritebiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/favourite/favouritemodel"
)

// giao dien danh cho storage thuc hien xoa
type DeleteFavouriteStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*favouritemodel.Favourite, error)
	Delete(context context.Context, userId int, foodId int) error
}

// Cau truc cua doi tuong xoa nha hang
type deleteFavouriteBiz struct {
	store     DeleteFavouriteStore
	requester common.Requester
}

// Khoi tai biz delete
func NewDeleteCartBiz(store DeleteFavouriteStore, requester common.Requester) *deleteFavouriteBiz {
	return &deleteFavouriteBiz{store: store, requester: requester}
}

// Ham Thuc Thi Xoa Nha Hang
func (biz *deleteFavouriteBiz) DeleteFavourite(context context.Context, userId int, foodId int) error {
	oldData, err := biz.store.
		FindDataWithCondition(context, map[string]interface{}{"food_id": foodId, "user_id": userId})

	if err != nil {
		return common.ErrEntityNotFound(favouritemodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(favouritemodel.EntityName, nil)
	}

	// if oldData.UserId != biz.requester.GetUserId() {
	// 	return common.ErrNoPermission(nil)
	// }

	if err := biz.store.Delete(context, userId, foodId); err != nil {
		return common.ErrCannotDeleteEntity(favouritemodel.EntityName, err)
	}

	return nil
}
