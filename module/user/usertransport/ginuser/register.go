package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/hasher"
	"food_delivery/module/user/userbusiness"
	"food_delivery/module/user/usermodel"
	"food_delivery/module/user/userstorage"

	"github.com/gin-gonic/gin"

	"net/http"
)

func Register(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		business := userbusiness.NewRegisterBusiness(store, md5)

		if err := business.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		// data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}

}
