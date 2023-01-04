package cardbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/card/cardmodel"
)

type CreateCardStore interface {
	Create(context context.Context, data *cardmodel.CardCreate) error
}

type createCardBiz struct {
	store CreateCardStore
}

func NewCreateCardBiz(store CreateCardStore) *createCardBiz {
	return &createCardBiz{store: store}
}

func (biz *createCardBiz) CreateCard(context context.Context, data *cardmodel.CardCreate) error {

	if err := biz.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(cardmodel.EntityName, err)
	}

	return nil
}
