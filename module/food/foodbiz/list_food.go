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
func NewListRestaurantBiz(store ListFoodStore) *listFoodBiz {
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

	// ids := make([]int, len(result))

	// for i := range result {
	// 	ids[i] = result[i].Id
	// }

	// likeMap, err := biz.likeStore.GetRestaurantLikes(context, ids)

	// if err != nil {
	// 	log.Println("Something went wrong")
	// 	return result, nil
	// }

	// for i, item := range result {
	// 	result[i].LikedCount = likeMap[item.Id]
	// }

	return result, nil
}
