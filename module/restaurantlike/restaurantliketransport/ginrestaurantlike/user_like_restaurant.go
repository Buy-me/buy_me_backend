package ginrestaurantlike

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/restaurantlike/restaurantlikebiz"
	"food_delivery/module/restaurantlike/restaurantlikemodel"
	"food_delivery/module/restaurantlike/restaurantlikestorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//POST /v1/restaurants/:id/like

func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// uid, err := common.FromBase58(c.Param("id"))
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.RestaurantLike{
			RestaurantId: id,
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		// incStore := restaurantstorage.NewSQlStore(appCtx.GetMainDBConnection())

		biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store, appCtx.GetPubsub())

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
