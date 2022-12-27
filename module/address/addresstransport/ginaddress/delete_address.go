package ginaddress

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/address/addressbiz"
	"food_delivery/module/address/addressstorage"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteAddress(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		// uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := addressstorage.NewSQLStore(db)
		biz := addressbiz.NewDeleteAddressBiz(store, requester)

		if err := biz.DeleteAddress(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}

}
