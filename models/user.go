package models

type User struct {
	UserId      int64
	Name        string
	Email       string
	PhoneNumber string
	Password    string
}

type UserRequest struct {
	Name        string
	Email       string
	Password    string
	PhoneNumber string
}

type UserResponse struct {
	Name        string
	Email       string
	PhoneNumber string
}

type Favorite struct {
	UserId int64
	GameId int64
}
