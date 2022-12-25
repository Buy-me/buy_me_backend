package ginfood

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/food/foodbiz"
	"food_delivery/module/food/foodmodel"
	"food_delivery/module/food/foodstorage"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFoodRating(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data foodmodel.FoodRating

		if err := c.ShouldBind(&data); err != nil {

			panic(err)
		}

		data.UserId = requester.GetUserId()
		data.FoodId = id

		store := foodstorage.NewSQLStore(db)

		biz := foodbiz.NewCreateFoodRatingBiz(store)

		if err := biz.CreateFoodRating(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}

}
