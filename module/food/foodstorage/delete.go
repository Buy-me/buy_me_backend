package foodstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/foodmodel"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	if err := s.db.Table(foodmodel.Food{}.
		TableName()).
		Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
