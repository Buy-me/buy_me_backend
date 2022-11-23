package restaurantbiz

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"
	"log"
)

type LikeRestaurantStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

// Dinh Nghia giao dien cho storage thi hanh
type ListRestaurantStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

// Dinh nghia cau truc nghiep vu nha hang
type listRestaurantBiz struct {
	store     ListRestaurantStore
	likeStore LikeRestaurantStore
}

// Khoi tao doi tuong lay danh sach nha hang
func NewListRestaurantBiz(store ListRestaurantStore, likeStore LikeRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store, likeStore: likeStore}
}

// Dinh Nghia Chuc Nang Cho Transport Su Dung
func (biz *listRestaurantBiz) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.store.ListDataWithCondition(context, filter, paging, "User")

	if err != nil {
		return nil, err
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	likeMap, err := biz.likeStore.GetRestaurantLikes(context, ids)

	if err != nil {
		log.Println("Something went wrong")
		return result, nil
	}

	for i, item := range result {
		result[i].LikedCount = likeMap[item.Id]
	}

	return result, nil
}
