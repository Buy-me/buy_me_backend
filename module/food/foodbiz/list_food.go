package foodbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/foodmodel"
)

// type LikeRestaurantStore interface {
// 	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
// }

// Dinh Nghia giao dien cho storage thi hanh
type ListFoodStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *foodmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]foodmodel.Food, error)
}

// Dinh nghia cau truc nghiep vu nha hang
type listFoodBiz struct {
	store ListFoodStore
	// likeStore LikeRestaurantStore
}

// Khoi tao doi tuong lay danh sach nha hang
func NewListFoodBiz(store ListFoodStore) *listFoodBiz {
	return &listFoodBiz{store: store}
}

// Dinh Nghia Chuc Nang Cho Transport Su Dung
func (biz *listFoodBiz) ListFood(
	context context.Context,
	filter *foodmodel.Filter,
	paging *common.Paging,
) ([]foodmodel.Food, error) {

	result, err := biz.store.ListDataWithCondition(context, filter, paging, "User")

	if err != nil {
		return nil, err
	}

	return result, nil
}
