package ginticket

import (
	"food_delivery/common"
	"food_delivery/component/appctx"

	"food_delivery/module/ticket/ticketbiz"
	"food_delivery/module/ticket/ticketmodel"
	"food_delivery/module/ticket/ticketstorage"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTicket(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		// requester := c.MustGet(common.CurrentUser).(common.Requester)
		log.Println("Get DB Connect Successfully")
		var data ticketmodel.TicketCreate

		if err := c.ShouldBind(&data); err != nil {
			// c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			// return
			panic(err)
		}

	
		
		// data.UserId = requester.GetUserId()

		store := ticketstorage.NewSQLStore(db)

		biz := ticketbiz.NewCreateTicketBiz(store)

		if err := biz.CreateTicket(c.Request.Context(), &data); err != nil {
			// c.JSON(http.StatusBadRequest, err)
			// return
			panic(err)
		}
		log.Println("Create oke")

		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}

}
