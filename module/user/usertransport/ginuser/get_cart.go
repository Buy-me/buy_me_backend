package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/cart/cartmodel"
	"food_delivery/module/user/userbusiness"
	"food_delivery/module/user/userstorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCart(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := userstorage.NewSQLStore(db)
		biz := userbusiness.NewGetCartBiz(store, requester)

		data, err := biz.GetCart(c.Request.Context(), requester.GetUserId())

		if err != nil {
			panic(err)
		}

		//WTF
		listFood := data.ListFoodInCart
		arrIds := make([]int, len(listFood))

		for index, item := range listFood {
			arrIds[index] = item.Id
		}

		var carts []cartmodel.Cart
		err = db.Where("food_id in (?) and status in (1)", arrIds).Find(&carts).Error

		for index := range carts {
			carts[index].Food = &listFood[index]
		}

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(carts))
	}

}
