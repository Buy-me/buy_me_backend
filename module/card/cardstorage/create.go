package cardstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/card/cardmodel"
	"log"
)

func (s *sqlStore) Create(context context.Context, data *cardmodel.CardCreate) error {

	var oldCard *cardmodel.Card
	s.db.Where("user_id = ? and type_card = ?", data.UserId, data.TypeCard).First(&oldCard)

	log.Println("oldCard", oldCard)
	if oldCard.UserId != 0 {
		if err := s.db.Table(cardmodel.Card{}.TableName()).
			Where("user_id= ? and type_card = ?", oldCard.UserId, oldCard.TypeCard).
			Updates(data).Error; err != nil {
			return common.ErrDB(err)
		}
	} else {
		if err := s.db.Create(&data).Error; err != nil {
			return common.ErrDB(err)
		}
	}

	return nil
}
