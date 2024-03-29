package userstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"
	"strings"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *usermodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]usermodel.User, error) {

	var result []usermodel.User

	//or (1,2,3)
	db := s.db.Table(usermodel.User{}.TableName()).Where("status in (1)")

	if filter != nil {
		if filter.Search != "" {
			db = db.Where("email LIKE ?", "%"+strings.Trim(filter.Search, " ")+"%")
		}

		if filter.Sort != "" {
			db = db.Order(filter.Sort)
		} else {
			db = db.Order("id desc")
		}
	} else {
		db = db.Order("id desc")
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit

		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
