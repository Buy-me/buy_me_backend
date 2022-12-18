package ginrestaurantlike

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/restaurantlike/restaurantlikebiz"
	"food_delivery/module/restaurantlike/restaurantlikestorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserUnlikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		// uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		// decStore := restaurantstorage.NewSQlStore(appCtx.GetMainDBConnection())

		biz := restaurantlikebiz.NewUserUnlikeRestaurantBiz(store, appCtx.GetPubsub())
		// biz := restaurantlikebiz.NewUserUnlikeRestaurantBiz(store, appCtx.GetPubsub())

		if err := biz.UnlikeRestaurant(c.Request.Context(), requester.GetUserId(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
