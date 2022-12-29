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

func ListUserOrder(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var filter ordermodel.Filter

		filter.UserId = requester.GetUserId()

		pagingData.Fulfill()

		var result []ordermodel.Order

		store := orderstorage.NewSQLStore(db)
		biz := orderbiz.NewListOrderBiz(store)

		result, err := biz.ListOrder(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}

}
