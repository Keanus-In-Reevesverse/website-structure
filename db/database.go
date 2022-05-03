package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnect() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err.Error())
	}

	db.Create("Table")
}
