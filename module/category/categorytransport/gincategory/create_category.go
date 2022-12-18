package gincategory

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/category/categorybiz"
	"food_delivery/module/category/categorymodel"
	"food_delivery/module/category/categorystorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		// requester := c.MustGet(common.CurrentUser).(common.Requester)
		var data categorymodel.CategoryCreate

		if err := c.ShouldBind(&data); err != nil {
			// c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			// return
			panic(err)
		}

		store := categorystorage.NewSQLStore(db)

		biz := categorybiz.NewCreateCategoryBiz(store)

		if err := biz.CreateCategory(c.Request.Context(), &data); err != nil {
			// c.JSON(http.StatusBadRequest, err)
			// return
			panic(err)
		}

		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}

}
