package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/hasher"
	"food_delivery/module/user/userbusiness"
	"food_delivery/module/user/usermodel"
	"food_delivery/module/user/userstorage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ChangePassword(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		user := c.MustGet(common.CurrentUser).(*usermodel.User)
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var dataChange usermodel.UserChangePassword
		if err := c.ShouldBind(&dataChange); err != nil {

			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbusiness.NewChangePasswordBiz(store, requester, md5)

		passHashed := md5.Hash(dataChange.OldPassword + user.Salt)
		if user.Password != passHashed {
			panic(usermodel.ErrPasswordInvalid)
		}

		if dataChange.NewPassword == dataChange.OldPassword {
			panic(common.NewFullErrorResponse(400, nil, "The old password and new password are the same", "The old password and new password are the same", "ChangePassword"))
		}

		var data usermodel.UserUpdate
		data.Password = dataChange.NewPassword
		err := biz.ChangePassword(c.Request.Context(), data, requester.GetUserId())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
