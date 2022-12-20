package ginticket

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/ticket/ticketbiz"
	"food_delivery/module/ticket/ticketstorage"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func PaymentTicket(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := ticketstorage.NewSQLStore(db)
		biz := ticketbiz.NewUpdateTicketBiz(store)

		state := "success"

		err = biz.UpdateState(c.Request.Context(), state, id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
