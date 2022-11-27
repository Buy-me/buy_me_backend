package ginrestaurantlike

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/restaurantlike/restaurantlikebiz"
	"food_delivery/module/restaurantlike/restaurantlikemodel"
	"food_delivery/module/restaurantlike/restaurantlikestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUsersLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("id"))

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
			// return
		}

		paging.Fulfill()

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.NewListUsersLikeRestaurantBiz(store)

		result, err := biz.ListUsersLikeRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
			// return
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))

	}

}
