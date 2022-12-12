package ginticket

import (
	"food_delivery/common"
	"food_delivery/component/appctx"

	"food_delivery/module/ticket/ticketbiz"
	"food_delivery/module/ticket/ticketstorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTicket(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		// id, err := strconv.Atoi(c.Param("id"))

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := ticketstorage.NewSQLStore(db)
		biz := ticketbiz.NewGetTicketBiz(store)

		data, err := biz.GetTicket(c.Request.Context(), int(uid.GetLocalID()))

		data.Mask(false)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}

}
