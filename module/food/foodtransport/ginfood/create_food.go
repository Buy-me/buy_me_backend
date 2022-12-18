package ginfood

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/food/foodbiz"
	"food_delivery/module/food/foodmodel"
	"food_delivery/module/food/foodstorage"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		// requester := c.MustGet(common.CurrentUser).(common.Requester)
		log.Println("Get DB Connect Successfully")
		var data foodmodel.FoodCreate

		if err := c.ShouldBind(&data); err != nil {
			// c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			// return
			panic(err)
		}

		log.Println("Parse data", data)
		// data.UserId = requester.GetUserId()

		store := foodstorage.NewSQLStore(db)

		biz := foodbiz.NewCreateFoodBiz(store)

		if err := biz.CreateFood(c.Request.Context(), &data); err != nil {
			// c.JSON(http.StatusBadRequest, err)
			// return
			panic(err)
		}
		log.Println("Create oke")

		// data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}

}
