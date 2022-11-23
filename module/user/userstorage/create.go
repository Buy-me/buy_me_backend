package userstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"
)

func (s *sqlStore) CreateUser(ctx context.Context, createUserData *usermodel.UserCreate) error {
	db := s.db.Begin()

	if err := db.Table(usermodel.UserCreate{}.TableName()).Create(&createUserData).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
