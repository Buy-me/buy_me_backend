package foodbiz

import (
	"context"
	"food_delivery/module/food/foodmodel"
)

// type LikeRestaurantStore interface {
// 	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
// }

// Dinh Nghia giao dien cho storage thi hanh
type ListFoodRatingStore interface {
	ListFoodRating(context context.Context, foodId int, moreKeys ...string) ([]foodmodel.FoodRating, error)
}

// Dinh nghia cau truc nghiep vu nha hang
type listFoodRatingBiz struct {
	store ListFoodRatingStore
	// likeStore LikeRestaurantStore
}

// Khoi tao doi tuong lay danh sach nha hang
func NewListFoodRatingBiz(store ListFoodRatingStore) *listFoodRatingBiz {
	return &listFoodRatingBiz{store: store}
}

// Dinh Nghia Chuc Nang Cho Transport Su Dung
func (biz *listFoodRatingBiz) ListFoodRating(context context.Context, foodId int) ([]foodmodel.FoodRating, error) {

	result, err := biz.store.ListFoodRating(context, foodId)

	if err != nil {
		return nil, err
	}

	return result, nil
}
