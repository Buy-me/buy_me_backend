package ginfavourite

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/favourite/favouritebiz"
	"food_delivery/module/favourite/favouritestorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteFavourite(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		type favouriteUpdate struct {
			FoodId int `json:"food_id"`
		}

		var data favouriteUpdate

		err := c.ShouldBind(&data)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := favouritestorage.NewSQLStore(db)
		biz := favouritebiz.NewDeleteCartBiz(store, requester)

		if err := biz.DeleteFavourite(c.Request.Context(), requester.GetUserId(), data.FoodId); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}

}
