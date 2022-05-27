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

func CreateUser(userCreated *models.UserRequest) *models.User {
	var user models.User

	user.Name = userCreated.Name
	user.Email = userCreated.Email
	user.PhoneNumber = userCreated.PhoneNumber
	user.Password = userCreated.Password

	DB.Table("USER").Create(&user)

	return &user
}
