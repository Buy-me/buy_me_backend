package foodstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/foodmodel"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *foodmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]foodmodel.Food, error) {

	var result []foodmodel.Food

	//or (1,2,3)
	db := s.db.Table(foodmodel.Food{}.TableName()).Where("status in (1)")

	// if filter != nil {
	// 	if filter.OwnerId > 0 {
	// 		db = db.Where("user_id = ?", filter.OwnerId)
	// 	}
	// }

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	// for i := range moreKeys {
	// 	db = db.Preload(moreKeys[i])
	// }

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
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
