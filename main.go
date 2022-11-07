package main

import (
	"awesomeProject1/component/appctx"
	"awesomeProject1/component/uploadprovider"
	"awesomeProject1/middleware"
	"awesomeProject1/module/restaurant/transport/ginnrestaurant"
	"awesomeProject1/module/upload/transport/ginupload"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"` //tag
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "root:@tcp(127.0.0.1:3306)/delivery?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := os.Getenv("MYSQL_CONN_STRING") // env

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	//secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appContext := appctx.NewAppContext(db, s3Provider)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Static("/static", "./static")

	//POST /restaurants
	v1 := r.Group("v1")

	v1.POST("/upload", ginupload.UploadImage(appContext))

	restaurants := v1.Group("/restaurants")
	restaurants.POST("", ginnrestaurant.CreateRestaurant(appContext))

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

		db.Where("id = ?", id).First(&data)

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

		db.Where("id = ?", id).Updates(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	//Delete Restaurants
	restaurants.DELETE("/:id", ginnrestaurant.DeleteRestaurant(appContext))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
