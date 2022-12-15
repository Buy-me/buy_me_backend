package ginorder

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/order/orderbiz"
	"food_delivery/module/order/ordermodel"
	"food_delivery/module/order/orderstorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOrder(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		// id, err := strconv.Atoi(c.Param("id"))

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data ordermodel.OrderUpdate

		err = c.ShouldBind(&data)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := orderstorage.NewSQLStore(db)
		biz := orderbiz.NewUpdateOrderBiz(store, requester)

		err = biz.UpdateFood(c.Request.Context(), data, int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
