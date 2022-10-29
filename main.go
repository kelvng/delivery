package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "root:@tcp(127.0.0.1:3306)/delivery?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(db, err)
}
