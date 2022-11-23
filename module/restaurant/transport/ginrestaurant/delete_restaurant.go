package ginrestaurant

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurantbiz "food_delivery/module/restaurant/biz"
	restaurantstorage "food_delivery/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		// id, err := strconv.Atoi(c.Param("id"))

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			// c.JSON(http.StatusBadRequest, gin.H{
			// 	"error": err.Error(),
			// })

			// return
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQlStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store, requester)

		if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			// c.JSON(http.StatusBadRequest, gin.H{
			// 	"error": err.Error(),
			// })
			// return
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}

}
