package gincard

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/card/cardbiz"
	"food_delivery/module/card/cardmodel"
	"food_delivery/module/card/cardstorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCard(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		var data cardmodel.CardCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		data.UserId = requester.GetUserId()

		store := cardstorage.NewSQLStore(db)

		biz := cardbiz.NewCreateCardBiz(store)

		if err := biz.CreateCard(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
