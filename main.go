package main

import (
	"awesomeProject1/component/appctx"
	"awesomeProject1/component/uploadprovider"
	"awesomeProject1/middleware"
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

	dsn := os.Getenv("bc8c4cf9cd5307:de804454@us-cdbr-east-06.cleardb.net/heroku_1e50603699e4adf?reconnect=true") // env

	//s3BucketName := os.Getenv("S3BucketName")
	//s3Region := os.Getenv("S3Region")
	//s3APIKey := os.Getenv("S3APIKey")
	//s3SecretKey := os.Getenv("S3SecretKey")
	//s3Domain := os.Getenv("S3Domain")
	//secretKey := os.Getenv("SYSTEM_SECRET")

	s3BucketName := os.Getenv("delivery-golang")
	s3Region := os.Getenv("ap-southeast-1")
	s3APIKey := os.Getenv("AKIA2LX3UTKASMGBRKNV")
	s3SecretKey := os.Getenv("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")
	s3Domain := os.Getenv("https://d2ygp6qy0o4yxw.cloudfront.net")
	secretKey := os.Getenv("MaiKhaAi")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appContext := appctx.NewAppContext(db, s3Provider, secretKey)

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

	setupRoute(appContext, v1)
	setupAdminRoute(appContext, v1)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
