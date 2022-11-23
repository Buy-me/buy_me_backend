package middleware

import (
	"errors"
	"food_delivery/common"
	"food_delivery/component/appctx"

	"github.com/gin-gonic/gin"
)

func RoleRequired(appCtx appctx.AppContext, allowRoles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		isAllowed := false

		for _, item := range allowRoles {
			if u.GetRole() == item {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			panic(common.ErrNoPermission(errors.New("invalid role user")))
		}

		c.Next()

	}
}
