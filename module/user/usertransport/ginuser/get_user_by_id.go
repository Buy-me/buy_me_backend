package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/user/userbusiness"
	"food_delivery/module/user/userstorage"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		biz := userbusiness.NewGetUserBiz(store, requester)

		data, err := biz.GetUser(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}

}
