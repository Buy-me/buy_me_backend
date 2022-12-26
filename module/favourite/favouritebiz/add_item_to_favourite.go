package favouritebiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/favourite/favouritemodel"
)

type CreateFavouriteStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*favouritemodel.Favourite, error)
	Create(context context.Context, data *favouritemodel.FavouriteCreate) error
}

type createFavouriteBiz struct {
	store CreateFavouriteStore
}

func NewCreateFavouriteBiz(store CreateFavouriteStore) *createFavouriteBiz {
	return &createFavouriteBiz{store: store}
}

func (biz *createFavouriteBiz) CreateFavourite(context context.Context, data *favouritemodel.FavouriteCreate) error {
	oldData, _ := biz.store.
		FindDataWithCondition(context, map[string]interface{}{"food_id": data.FoodId, "user_id": data.UserId})

	if oldData != nil {
		return common.ErrEntityExisted(favouritemodel.EntityName+" items", nil)
	}

	if err := biz.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(favouritemodel.EntityName, err)
	}

	return nil
}
