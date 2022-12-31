package main

import (
	"awesomeProject1/component/appctx"
	"awesomeProject1/component/uploadprovider"
	"awesomeProject1/middleware"
	"awesomeProject1/pubsub/localpd"
	"awesomeProject1/skio"
	"awesomeProject1/subscriber"
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
	dsn := os.Getenv("MYSQL_CONN_STRING") // env

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")

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
	ps := localpd.NewPubSub()
	appContext := appctx.NewAppContext(db, s3Provider, secretKey, ps)

	//set up subscriber

	//subscriber.Setup(appContext, context.Background())

	_ = subscriber.NewEngine(appContext).Start()

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

	rtEngine := skio.NewEngine()
	appContext.SetRealtimeEngine(rtEngine)

	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
