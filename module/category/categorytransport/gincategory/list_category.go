package gincategory

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/category/categorybiz"
	"food_delivery/module/category/categorymodel"
	"food_delivery/module/category/categorystorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var filter categorymodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var result []categorymodel.Category

		store := categorystorage.NewSQLStore(db)
		biz := categorybiz.NewListCategoryBiz(store)

		result, err := biz.ListCategory(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}

}
