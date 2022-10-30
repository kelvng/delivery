package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "root:@tcp(127.0.0.1:3306)/delivery?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := os.Getenv("MYSQL_CONN_STRING") // env
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// Insert Database
	newRestaurant := Restaurant{Name: "Khang", Addr: "Ben Tre"}

	if err := db.Create(&newRestaurant).Error; err != nil {
		log.Println(err)
	} else {
		log.Println("New id:", newRestaurant.Id)
	}

	// Select Database
	var myRestaurant Restaurant

	if err := db.Where("id = ?", 9).First(&myRestaurant).Error; err != nil {
		log.Println(err)
	} else {
		log.Println(myRestaurant)
	}

	// Update Database
	newName := ""
	updateData := RestaurantUpdate{Name: &newName}

	if err := db.Where("id = ?", 9).Updates(&updateData).Error; err != nil {
		log.Println(err)
	} else {
		log.Println(myRestaurant)
	}

	// Delete Database
	if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 10).Delete(&myRestaurant).Error; err != nil {
		log.Println(err)
	} else {
		log.Println(myRestaurant)
	}
}
