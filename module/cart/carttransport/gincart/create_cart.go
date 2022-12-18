package gincart

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/cart/cartbiz"
	"food_delivery/module/cart/cartmodel"
	"food_delivery/module/cart/cartstorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		// requester := c.MustGet(common.CurrentUser).(common.Requester)
		var data cartmodel.CartCreate

		if err := c.ShouldBind(&data); err != nil {
			// c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			// return
			panic(err)
		}

		store := cartstorage.NewSQLStore(db)

		biz := cartbiz.NewCreateCartBiz(store)

		if err := biz.CreateCart(c.Request.Context(), &data); err != nil {
			// c.JSON(http.StatusBadRequest, err)
			// return
			panic(err)
		}

		// data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}

}
