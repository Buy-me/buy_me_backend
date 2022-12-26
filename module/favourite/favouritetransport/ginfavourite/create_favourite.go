package ginfavourite

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/favourite/favouritebiz"
	"food_delivery/module/favourite/favouritemodel"
	"food_delivery/module/favourite/favouritestorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFavourite(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		var data favouritemodel.FavouriteCreate

		if err := c.ShouldBind(&data); err != nil {
			// c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			// return
			panic(err)
		}

		data.UserId = requester.GetUserId()

		store := favouritestorage.NewSQLStore(db)

		biz := favouritebiz.NewCreateFavouriteBiz(store)

		if err := biz.CreateFavourite(c.Request.Context(), &data); err != nil {
			// c.JSON(http.StatusBadRequest, err)
			// return
			panic(err)
		}

		// data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}

}
