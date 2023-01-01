package ginfood

import (
	"fmt"
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/food/foodbiz"
	"food_delivery/module/food/foodmodel"
	"food_delivery/module/food/foodstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListFoodRating(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var result []foodmodel.FoodRating

		store := foodstorage.NewSQLStore(db)
		biz := foodbiz.NewListFoodRatingBiz(store)

		result, err = biz.ListFoodRating(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		//WTF

		arrIds := make([]int, len(result))

		for index, item := range result {
			arrIds[index] = item.UserId
		}

		var users []foodmodel.User
		err = db.Where("id in (?) and status in (1)", arrIds).Find(&users).Error

		if err != nil {
			panic(err)
		}
		// arrResultIds := make([]int, len(users))
		mapResults := make(map[int]*foodmodel.User)
		for index := range users {
			mapResults[users[index].Id] = &users[index]
		}

		fmt.Println("users", users)
		fmt.Println("mapResults", mapResults)

		for index := range result {
			result[index].User = mapResults[result[index].UserId]
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}

}
