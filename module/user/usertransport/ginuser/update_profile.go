package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/user/userbusiness"
	"food_delivery/module/user/usermodel"
	"food_delivery/module/user/userstorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProfile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		// uid, err := common.FromBase58(c.Param("id"))

		var data usermodel.UserUpdate

		err := c.ShouldBind(&data)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		biz := userbusiness.NewUpdateProfileBiz(store, requester)

		err = biz.UpdateProfile(c.Request.Context(), data, requester.GetUserId())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
