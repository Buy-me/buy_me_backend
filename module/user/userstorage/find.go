package userstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(
	context context.Context,
	conditions map[string]interface{},
	moreInfo ...string,
) (*usermodel.User, error) {

	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
