package ginfood

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/food/foodbiz"
	"food_delivery/module/food/foodstorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		// id, err := strconv.Atoi(c.Param("id"))

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodstorage.NewSQLStore(db)
		biz := foodbiz.NewGetFoodBiz(store, requester)

		data, err := biz.GetFood(c.Request.Context(), int(uid.GetLocalID()))

		data.Mask(false)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}

}
