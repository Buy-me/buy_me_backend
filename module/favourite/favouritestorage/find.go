package favouritestorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/favourite/favouritemodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*favouritemodel.Favourite, error) {

	var data favouritemodel.Favourite

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
