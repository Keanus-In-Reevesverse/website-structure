package models

import "gopkg.in/validator.v2"

type Login struct {
	Email    string `json:"email" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
}

func LoginValidator(login *Login) error {
	if err := validator.Validate(login); err != nil {
		return err
	}

	return nil
}
