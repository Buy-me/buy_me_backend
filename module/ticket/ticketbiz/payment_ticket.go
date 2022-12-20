package ticketbiz

import (
	"context"
)

type UpdateTicketStore interface {
	UpdateState(ctx context.Context, state string, id int) error
}

type updateTicketBiz struct {
	store UpdateTicketStore
}

func NewUpdateTicketBiz(store UpdateTicketStore) *updateTicketBiz {
	return &updateTicketBiz{store: store}
}

func (biz *updateTicketBiz) UpdateState(context context.Context, state string, id int) error {
	// oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	err := biz.store.UpdateState(context, state, id)

	if err != nil {
		return err
	}

	return nil
}
