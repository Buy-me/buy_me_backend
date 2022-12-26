package addressstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/address/addressmodel"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	if err := s.db.Table(addressmodel.Address{}.
		TableName()).
		Where("id = ?", id).
		Delete(&addressmodel.Address{}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
