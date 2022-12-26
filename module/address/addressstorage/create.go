package addressstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/address/addressmodel"
)

func (s *sqlStore) Create(context context.Context, data *addressmodel.AddressCreate) error {

	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
