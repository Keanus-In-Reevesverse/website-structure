package database

import (
	"log"

	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnect() {
	conn := "admin:secret@tcp(127.0.0.1:3306)/buscajogos?charset=utf8mb4&parseTime=True&loc=Local"
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
