package userstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"
)

func (s *sqlStore) Update(ctx context.Context, data usermodel.UserUpdate, id int) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
