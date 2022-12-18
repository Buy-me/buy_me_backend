package cartstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/cart/cartmodel"
)

func (s *sqlStore) Create(context context.Context, data *cartmodel.CartCreate) error {

	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
