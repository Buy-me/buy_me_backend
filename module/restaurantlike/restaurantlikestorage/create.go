package restaurantlikestorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurantlike/restaurantlikemodel"
	"time"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error {
	now := time.Now().UTC()
	data.CreatedAt = &now
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
