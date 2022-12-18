package categorybiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/category/categorymodel"
)

// type LikeRestaurantStore interface {
// 	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
// }

// Dinh Nghia giao dien cho storage thi hanh
type ListCategoryStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *categorymodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]categorymodel.Category, error)
}

// Dinh nghia cau truc nghiep vu nha hang
type listCategoryBiz struct {
	store ListCategoryStore
	// likeStore LikeRestaurantStore
}

// Khoi tao doi tuong lay danh sach nha hang
func NewListCategoryBiz(store ListCategoryStore) *listCategoryBiz {
	return &listCategoryBiz{store: store}
}

// Dinh Nghia Chuc Nang Cho Transport Su Dung
func (biz *listCategoryBiz) ListCategory(
	context context.Context,
	filter *categorymodel.Filter,
	paging *common.Paging,
) ([]categorymodel.Category, error) {

	result, err := biz.store.ListDataWithCondition(context, filter, paging)

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
