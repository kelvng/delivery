package main

import (
	"awesomeProject1/component/appctx"
	"awesomeProject1/middleware"
	"awesomeProject1/module/restaurant/transport/ginnrestaurant"
	"awesomeProject1/module/upload/transport/ginupload"
	"awesomeProject1/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func setupRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {

	v1.POST("/upload", ginupload.UploadImage(appContext))

	v1.POST("/register", ginuser.Register(appContext))

	v1.POST("/authenticate", ginuser.Login(appContext))

	v1.GET("/profile", middleware.RoleRRequired(appContext), ginuser.Profile(appContext))

	restaurants := v1.Group("/restaurants", middleware.RoleRRequired(appContext))

	{
		restaurants.POST("", ginnrestaurant.CreateRestaurant(appContext))
	}

	//GET ID
	restaurants.GET("/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data Restaurant

		appContext.GetMaiDBConnection().Where("id = ?", id).First(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	// List Restaurants
	restaurants.GET("", ginnrestaurant.ListRestaurant(appContext))

	//Update Restaurants
	restaurants.PATCH("/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		appContext.GetMaiDBConnection().Where("id = ?", id).Updates(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	//Delete Restaurants
	restaurants.DELETE("/:id", ginnrestaurant.DeleteRestaurant(appContext))

}
