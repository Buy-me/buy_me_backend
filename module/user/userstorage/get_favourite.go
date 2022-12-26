package userstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"

	"gorm.io/gorm"
)

func (s *sqlStore) GetFavourite(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*usermodel.User, error) {

	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var user usermodel.User

	if err := db.Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
