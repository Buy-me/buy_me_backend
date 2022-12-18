package ginorder

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/order/orderbiz"
	"food_delivery/module/order/ordermodel"
	"food_delivery/module/order/orderstorage"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data ordermodel.OrderCreate

		if err := c.ShouldBind(&data); err != nil {
			// c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			// return
			panic(err)
		}

		log.Println(data.Items)
		log.Println(data.TotalPrice)

		data.UserId = requester.GetUserId()

		store := orderstorage.NewSQLStore(db)

		biz := orderbiz.NewCreateTicketBiz(store)

		if err := biz.CreateOrder(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		// data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))

	}

}
