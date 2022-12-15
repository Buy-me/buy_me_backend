package main

import (
	"food_delivery/component/appctx"
	"food_delivery/middleware"
	"food_delivery/module/food/foodtransport/ginfood"
	"food_delivery/module/order/ordertransport/ginorder"
	"food_delivery/module/restaurant/transport/ginrestaurant"
	"food_delivery/module/restaurantlike/restaurantliketransport/ginrestaurantlike"
	"food_delivery/module/ticket/tickettransport/ginticket"
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
	{

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
		restaurant.POST("/:id/like", ginrestaurantlike.UserLikeRestaurant(appContext))
		restaurant.DELETE("/:id/unlike", ginrestaurantlike.UserUnlikeRestaurant(appContext))
		restaurant.GET("/:id/liked-users", ginrestaurantlike.ListUsersLikeRestaurant(appContext))
	}

	food := v1.Group("/foods", middleware.RequiredAuth(appContext))
	{
		food.POST("/", ginfood.CreateFood(appContext))
		food.GET("/", ginfood.ListFood(appContext))
		food.GET("/:id", ginfood.GetFood(appContext))
		food.DELETE("/:id", ginfood.DeleteFood(appContext))
		food.PATCH("/:id", ginfood.UpdateFood(appContext))
	}

	ticket := v1.Group("/tickets")
	{
		ticket.POST("/", ginticket.CreateTicket(appContext))
		ticket.GET("/", ginticket.ListTicket(appContext))
		ticket.GET("/:id", ginticket.GetTicket(appContext))
	}

	order := v1.Group("/orders", middleware.RequiredAuth(appContext))
	{
		order.POST("/", ginorder.CreateOrder(appContext))
		order.GET("/", ginorder.ListOrder(appContext))
		order.GET("/:id", ginorder.GetOrder(appContext))
		order.DELETE("/:id", ginorder.DeleteOrder(appContext))
		order.PATCH("/:id", ginorder.UpdateOrder(appContext))
	}

	// v1/restaurants/:id/liked-users
}
