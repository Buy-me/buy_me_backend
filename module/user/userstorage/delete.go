package userstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	if err := s.db.Table(usermodel.User{}.
		TableName()).
		Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
