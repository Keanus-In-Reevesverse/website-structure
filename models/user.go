package models

import "gopkg.in/validator.v2"

type User struct {
	UserId      int64  `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserRequest struct {
	Name        string `json:"name" validate:"max=60, nonzero"`
	Email       string `json:"email" validate:"max=60, nonzero"`
	PhoneNumber string `json:"phone_number" validate:"len=12, nonzero"`
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
