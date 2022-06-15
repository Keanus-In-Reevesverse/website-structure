package models

import "gopkg.in/validator.v2"

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
