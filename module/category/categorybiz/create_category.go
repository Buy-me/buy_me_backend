package categorybiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/category/categorymodel"
)

type CreateCategoryStore interface {
	Create(context context.Context, data *categorymodel.CategoryCreate) error
}

type createCategoryBiz struct {
	store CreateCategoryStore
}

func NewCreateCategoryBiz(store CreateCategoryStore) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateCategory(context context.Context, data *categorymodel.CategoryCreate) error {

	if err := biz.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}

	return nil
}
