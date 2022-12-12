package ginticket

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/ticket/ticketbiz"
	"food_delivery/module/ticket/ticketmodel"
	"food_delivery/module/ticket/ticketstorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTicket(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var filter ticketmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var result []ticketmodel.Ticket

		store := ticketstorage.NewSQLStore(db)
		biz := ticketbiz.NewListTicketBiz(store)

		result, err := biz.ListTicket(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}

}
