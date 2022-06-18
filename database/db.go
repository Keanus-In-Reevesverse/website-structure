package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnect() {
	admin := os.Getenv("DB_ADMIN")
	adminPass := os.Getenv("DB_ADMIN_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", admin, adminPass, host, port, name)
	DB, err = gorm.Open(mysql.Open(conn))

	if err != nil {
		log.Panic(err.Error())
	}

}

func UserOps(user *models.User, operation string) *models.User {
	if operation == "create" {
		DB.Table("USER").Create(&user)
	}

	if operation == "edit" {
		DB.Table("USER").Where("user_id = ?", &user.UserId).Updates(&user)
	}

	if operation == "delete" {
		DB.Table("USER").Where("user_id = ?", &user.UserId).Unscoped().Delete(&user)
	}

	return user
}

func AlertOps(alert models.Alert, operation string) models.Alert {
	if operation == "create" {
		DB.Table("USER").Create(&alert)
	}

	if operation == "edit" {
		DB.Table("USER").Where("game_id = ?", &alert.UserId).Updates(&alert)
	}

	if operation == "delete" {
		DB.Table("USER").Where("user_id = ?", &alert.UserId).Unscoped().Delete(&alert)
	}

	return alert
}
