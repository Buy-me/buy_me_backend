package ticketbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/ticket/ticketmodel"
)

// giao dien danh cho storage thuc hien xoa
type GetTicketStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*ticketmodel.Ticket, error)
}

// Cau truc cua doi tuong xoa nha hang
type getTicketBiz struct {
	store GetTicketStore
}

// Khoi tai biz delete
func NewGetTicketBiz(store GetTicketStore) *getTicketBiz {
	return &getTicketBiz{store: store}
}

func (biz *getTicketBiz) GetTicket(context context.Context, id int) (*ticketmodel.Ticket, error) {
	// oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	ticket, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrEntityNotFound(ticketmodel.EntityName, err)
	}

	if ticket.Status == 0 {
		return nil, common.ErrEntityDeleted(ticketmodel.EntityName, nil)
	}

	return ticket, nil
}
