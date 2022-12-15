package orderstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/order/ordermodel"
)

func (s *sqlStore) Create(context context.Context, data *ordermodel.OrderCreate) error {

	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
