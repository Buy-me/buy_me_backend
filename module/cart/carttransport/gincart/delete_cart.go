package gincart

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/cart/cartbiz"
	"food_delivery/module/cart/cartstorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCart(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		type cartUpdate struct {
			FoodId int `json:"food_id"`
		}

		var data cartUpdate

		err := c.ShouldBind(&data)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := cartstorage.NewSQLStore(db)
		biz := cartbiz.NewDeleteCartBiz(store, requester)

		if err := biz.DeleteCart(c.Request.Context(), requester.GetUserId(), data.FoodId); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}

}
