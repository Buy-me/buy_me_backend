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

func DeleteUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		// uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		biz := userbusiness.NewDeleteUserBiz(store, requester)

		if err := biz.DeleteFood(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
