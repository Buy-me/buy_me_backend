package ginorder

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/order/orderbiz"
	"food_delivery/module/order/ordermodel"
	"food_delivery/module/order/orderstorage"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUserOrderByUserId(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		userId, err := strconv.Atoi(c.Param("userId"))

		// requester := c.MustGet(common.CurrentUser).(common.Requester)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var filter ordermodel.Filter

		filter.UserId = userId

		pagingData.Fulfill()

		var result []ordermodel.Order

		store := orderstorage.NewSQLStore(db)
		biz := orderbiz.NewListOrderBiz(store)

		result, err = biz.ListOrder(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}

}
