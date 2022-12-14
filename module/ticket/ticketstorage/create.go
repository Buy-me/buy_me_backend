package ticketstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/ticket/ticketmodel"
)

func (s *sqlStore) Create(context context.Context, data *ticketmodel.TicketCreate) error {

	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
