package ticketstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/ticket/ticketmodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*ticketmodel.Ticket, error) {

	var data ticketmodel.Ticket
	db := s.db.Table(ticketmodel.Ticket{}.TableName())
	db = db.Preload("Travelers")

	if err := db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
