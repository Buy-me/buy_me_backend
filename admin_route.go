package main

import (
	"food_delivery/component/appctx"
	"food_delivery/middleware"
	"food_delivery/module/user/usertransport/ginuser"

	"github.com/gin-gonic/gin"
)

func setUpAdminRoutes(appContext appctx.AppContext, v1 *gin.RouterGroup) {

	admin := v1.Group("/admin",
		middleware.RequiredAuth(appContext),
		middleware.RoleRequired(appContext, "admin", "mod"),
	)
	{
		admin.GET("/profile", ginuser.Profile(appContext))
	}
}
