package favouritestorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/favourite/favouritemodel"
)

func (s *sqlStore) Create(context context.Context, data *favouritemodel.FavouriteCreate) error {

	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
