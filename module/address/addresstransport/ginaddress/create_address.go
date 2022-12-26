package ginaddress

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/address/addressbiz"
	"food_delivery/module/address/addressmodel"
	"food_delivery/module/address/addressstorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAddress(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		var data addressmodel.AddressCreate

		if err := c.ShouldBind(&data); err != nil {
			// c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			// return
			panic(err)
		}

		data.UserId = requester.GetUserId()

		store := addressstorage.NewSQLStore(db)

		biz := addressbiz.NewCreateAddressBiz(store)

		if err := biz.CreateAddress(c.Request.Context(), &data); err != nil {
			// c.JSON(http.StatusBadRequest, err)
			// return
			panic(err)
		}

		// data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}

}
