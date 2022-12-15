package orderstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/order/ordermodel"
)

func (s *sqlStore) Update(ctx context.Context, data ordermodel.OrderUpdate, id int) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
