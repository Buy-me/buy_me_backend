package ginfood

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/food/foodbiz"
	"food_delivery/module/food/foodmodel"
	"food_delivery/module/food/foodstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var filter foodmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var result []foodmodel.Food

		store := foodstorage.NewSQLStore(db)
		biz := foodbiz.NewListFoodBiz(store)

		result, err := biz.ListFood(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		// for i := range result {
		// 	result[i].Mask(false)
		// }
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}

}
