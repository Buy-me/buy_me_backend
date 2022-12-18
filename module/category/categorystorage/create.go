package categorystorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/category/categorymodel"
)

func (s *sqlStore) Create(context context.Context, data *categorymodel.CategoryCreate) error {

	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
