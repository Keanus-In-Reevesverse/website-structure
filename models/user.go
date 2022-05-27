package models

import "gopkg.in/validator.v2"

type User struct {
	UserId      int64
	Name        string
	Email       string
	PhoneNumber string
	Password    string
}

type UserRequest struct {
	Name        string `json:"name" validate:"max=60, nonzero"`
	Email       string `json:"email" validate:"max=60, nonzero"`
	PhoneNumber string `json:"phone_number" validate:"len=11, nonzero"`
	Password    string `json:"password" validate:"max=30, nonzero"`
}

func UserRequestValidator(userRequest *UserRequest) error {
	if err := validator.Validate(userRequest); err != nil {
		return err
	}

	return nil
}

type UserResponse struct {
	Name        string `json:"name" validate:"max=60, nonzero"`
	Email       string `json:"email" validate:"max=60, nonzero"`
	PhoneNumber string `json:"phone_number" validate:"len=12, nonzero"`
}

func UserResponseValidator(userResponse *UserResponse) error {
	if err := validator.Validate(userResponse); err != nil {
		return err
	}

	return nil
}

type Favorite struct {
	UserId int64 `json:"user_id" validate:"regexp=^[0-9]*$"`
	GameId int64 `json:"game_id" validate:"max=60, regexp=^[0-9]*$"`
}

func FavoriteValidator(favorite *Favorite) error {
	if err := validator.Validate(favorite); err != nil {
		return err
	}

	return nil
}
