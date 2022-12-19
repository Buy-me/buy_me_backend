package cartstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/cart/cartmodel"
)

func (s *sqlStore) Delete(context context.Context, userId int, foodId int) error {
	if err := s.db.Table(cartmodel.Cart{}.
		TableName()).
		Where("food_id = ? and user_id = ?", foodId, userId).
		Delete(&cartmodel.Cart{}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
