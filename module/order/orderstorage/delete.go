package orderstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/order/ordermodel"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	if err := s.db.Table(ordermodel.Order{}.
		TableName()).
		Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
