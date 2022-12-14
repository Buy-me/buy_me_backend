package ticketbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/ticket/ticketmodel"
)

// type LikeRestaurantStore interface {
// 	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
// }

// Dinh Nghia giao dien cho storage thi hanh
type ListTicketStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *ticketmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]ticketmodel.Ticket, error)
}

// Dinh nghia cau truc nghiep vu nha hang
type listTicketBiz struct {
	store ListTicketStore
	// likeStore LikeRestaurantStore
}

// Khoi tao doi tuong lay danh sach nha hang
func NewListTicketBiz(store ListTicketStore) *listTicketBiz {
	return &listTicketBiz{store: store}
}

// Dinh Nghia Chuc Nang Cho Transport Su Dung
func (biz *listTicketBiz) ListTicket(
	context context.Context,
	filter *ticketmodel.Filter,
	paging *common.Paging,
) ([]ticketmodel.Ticket, error) {

	result, err := biz.store.ListDataWithCondition(context, filter, paging, "Travelers")

	if err != nil {
		return nil, err
	}

	return result, nil
}
