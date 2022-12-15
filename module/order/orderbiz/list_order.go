package orderbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/order/ordermodel"
)

// type LikeRestaurantStore interface {
// 	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
// }

// Dinh Nghia giao dien cho storage thi hanh
type ListOrderStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *ordermodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]ordermodel.Order, error)
}

// Dinh nghia cau truc nghiep vu nha hang
type listOrderBiz struct {
	store ListOrderStore
	// likeStore LikeRestaurantStore
}

// Khoi tao doi tuong lay danh sach nha hang
func NewListOrderBiz(store ListOrderStore) *listOrderBiz {
	return &listOrderBiz{store: store}
}

// Dinh Nghia Chuc Nang Cho Transport Su Dung
func (biz *listOrderBiz) ListOrder(
	context context.Context,
	filter *ordermodel.Filter,
	paging *common.Paging,
) ([]ordermodel.Order, error) {

	result, err := biz.store.ListDataWithCondition(context, filter, paging, "Items")

	if err != nil {
		return nil, err
	}

	return result, nil
}
