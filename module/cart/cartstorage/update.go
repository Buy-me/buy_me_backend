package cartstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/cart/cartmodel"
)

// func (s *sqlStore) UpdateData(
// 	ctx context.Context,
// 	id int,
// 	data *restaurantmodel.RestaurantUpdate,
// ) error {
// 	data.PrepareForUpdate()
// 	db := s.db

// 	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
// 		return common.ErrDB(err)
// 	}
// 	return nil
// }

func (s *sqlStore) Update(ctx context.Context, quantity int, userId int, foodId int) error {
	db := s.db

	// if err := db.Table(foodmodel.Food{}.TableName()).Where("id = ?", id).
	// 	Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
	// 	return common.ErrDB(err)
	// }

	if err := db.Table(cartmodel.Cart{}.TableName()).
		Where("user_id= ? and food_id = ?", userId, foodId).
		Update("quantity", quantity).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}




// func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
// 	db := s.db

// 	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
// 		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
// 		return common.ErrDB(err)
// 	}

// 	return nil
// }

// func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
// 	db := s.db

// 	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
// 		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
// 		return common.ErrDB(err)
// 	}

// 	return nil
// }
