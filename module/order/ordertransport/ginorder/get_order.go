package ginorder

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/order/orderbiz"
	"food_delivery/module/order/orderstorage"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrder(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		// uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := orderstorage.NewSQLStore(db)

		biz := orderbiz.NewGetOrderBiz(store)

		data, err := biz.GetOrder(c.Request.Context(), id)

		// data.Mask(false)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}

}
