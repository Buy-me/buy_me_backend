package favouritestorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/cart/cartmodel"
	"food_delivery/module/favourite/favouritemodel"
)

func (s *sqlStore) Delete(context context.Context, userId int, foodId int) error {
	if err := s.db.Table(favouritemodel.Favourite{}.
		TableName()).
		Where("food_id = ? and user_id = ?", foodId, userId).
		Delete(&cartmodel.Cart{}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
