package restaurantlikestorage

import (
	"context"
	"fmt"
	"food_delivery/common"
	"food_delivery/module/restaurantlike/restaurantlikemodel"
	"strings"
	"time"

	"github.com/btcsuite/btcd/btcutil/base58"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	var listLike []restaurantlikemodel.LikedCount

	if err := s.db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).
		Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.RestaurantId] = item.LikeCount
	}

	return result, nil
}

func (s *sqlStore) GetUsersLikeRestaurant(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	// order *common.Order,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	var result []restaurantlikemodel.RestaurantLike

	db := s.db

	db = db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).Where(conditions)

	if filter != nil {
		if filter.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", filter.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Preload("User")

	if v := paging.FakeCursor; v != "" {
		// timeCreated, err := time.Parse(common.TimeFormat, string(base58.Decode(v)))
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))
		parts := strings.Split(timeCreated.String(), " ")

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", fmt.Sprintf("%v %v", parts[0], parts[1]))

	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	fmt.Println("==========", result, "================")

	usersLike := make([]common.SimpleUser, len(result))

	for i, item := range result {
		result[i].User.CreatedAt = item.CreatedAt
		result[i].User.UpdatedAt = nil
		usersLike[i] = *result[i].User

		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}
	return usersLike, nil
}
