package gincart

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/cart/cartbiz"
	"food_delivery/module/cart/cartstorage"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCart(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		type cartUpdate struct {
			Quantity int `json:"quantity"`
			FoodId   int `json:"food_id"`
		}

		var data cartUpdate

		err := c.ShouldBind(&data)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		log.Println("data", data)

		store := cartstorage.NewSQLStore(db)
		biz := cartbiz.NewUpdateCartBiz(store, requester)

		err = biz.UpdateCart(c.Request.Context(), data.Quantity, requester.GetUserId(), data.FoodId)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
