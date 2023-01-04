package orderstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/order/ordermodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*ordermodel.Order, error) {

	var data ordermodel.Order

	db := s.db.Table(ordermodel.Order{}.TableName())

	db = db.Preload("Items")

	if err := db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
