package ticketstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/ticket/ticketmodel"
	"log"
)

func (s *sqlStore) UpdateState(ctx context.Context, state string, id int) error {
	db := s.db

	// if err := db.Table(foodmodel.Food{}.TableName()).Where("id = ?", id).
	// 	Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
	// 	return common.ErrDB(err)
	// }

	log.Println("Come here")
	if err := db.Table(ticketmodel.Ticket{}.TableName()).Where("id = ?", id).Update("state", state).Error; err != nil {
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
