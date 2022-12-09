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

func UpdateFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		// id, err := strconv.Atoi(c.Param("id"))

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data foodmodel.FoodUpdate

		err = c.ShouldBind(&data)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodstorage.NewSQLStore(db)
		biz := foodbiz.NewUpdateFoodBiz(store, requester)

		err = biz.UpdateFood(c.Request.Context(), data, int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
