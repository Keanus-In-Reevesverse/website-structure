package models

type Login struct {
	Email    string
	Password string
}

type UserInput struct {
	Name        string
	Email       string
	Password    string
	PhoneNumber string
}

type UserOutput struct {
	Name        string
	Email       string
	PhoneNumber string
}
