package main

import (
	"food_delivery/component/appctx"
	"food_delivery/middleware"
	"food_delivery/module/restaurant/transport/ginrestaurant"
	"food_delivery/module/upload/transport/ginupload"
	"food_delivery/module/user/usertransport/ginuser"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func setUpRoutes(appContext appctx.AppContext, v1 *gin.RouterGroup) {

	db := appContext.GetMainDBConnection()
	v1.POST("/upload", ginupload.Upload(appContext))
	v1.POST("/register", ginuser.Register(appContext))
	v1.POST("/authenticate", ginuser.LoginHandler(appContext))
	v1.GET("/profile", middleware.RequiredAuth(appContext), ginuser.Profile(appContext))

	restaurant := v1.Group("/restaurants", middleware.RequiredAuth(appContext))

	restaurant.POST("/", ginrestaurant.CreateRestaurant(appContext))

	restaurant.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		var data Restaurant

		db.Where("id = ?", id).First(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	restaurant.GET("/", ginrestaurant.ListRestaurant(appContext))

	restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	restaurant.PATCH("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		var data RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		db.Where("id = ?", id).Updates(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": "Update Successfully",
		})
	})

}
