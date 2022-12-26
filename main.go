package main

import (
	"awesomeProject1/component/appctx"
	"awesomeProject1/component/uploadprovider"
	"awesomeProject1/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
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

	dsn := "bc8c4cf9cd5307:de804454@tcp(us-cdbr-east-06.cleardb.net)/heroku_1e50603699e4adf?charset=utf8mb4&parseTime=True&loc=Local" // env

	s3BucketName := "heroku_1e50603699e4adf"
	s3Region := "ap-southeast-1"
	s3APIKey := "AKIA2LX3UTKASMGBRKNV"
	s3SecretKey := "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"
	s3Domain := "https://d2ygp6qy0o4yxw.cloudfront.net"
	secretKey := "MaiKhaAi"
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appContext := appctx.NewAppContext(db, s3Provider, secretKey)

	r := gin.Default()

	r.Use(middleware.Recover(appContext))

	// Enable CORS for all routes
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	r.Use(cors.New(config))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Static("/static", "./static")

	//POST /restaurants
	v1 := r.Group("v1", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		// ...
	})

	setupRoute(appContext, v1)
	setupAdminRoute(appContext, v1)

	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
