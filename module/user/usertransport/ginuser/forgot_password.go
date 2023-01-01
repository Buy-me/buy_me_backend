package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/hasher"
	"food_delivery/module/user/userbusiness"
	"food_delivery/module/user/userstorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ForgotPassword(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		type ForgotPassword struct {
			Email string `json:"email"`
		}

		var body ForgotPassword

		if err := c.ShouldBind(&body); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		
		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbusiness.NewForgotPasswordBiz(store, md5)

		err := biz.ForgotPassword(c.Request.Context(), body.Email)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("A new password has been sent to your email"))
	}

}
