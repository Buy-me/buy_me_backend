package ginfood

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/food/foodbiz"
	"food_delivery/module/food/foodmodel"
	"food_delivery/module/food/foodstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListFoodRating(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var result []foodmodel.FoodRating

		store := foodstorage.NewSQLStore(db)
		biz := foodbiz.NewListFoodRatingBiz(store)

		result, err = biz.ListFoodRating(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}

}
