package entities

import "errors"

type User struct {
	ID          int    `json:"-" db:"id"`
	Username    string `json:"username" db:"username" binding:"required" `
	Name        string `json:"name" db:"name" binding:"required" `
	Surname     string `json:"surname" db:"surname" binding:"required"`
	Teacher     string `json:"teacher" db:"teacher" binding:"required" `
	Password    string `json:"password" db:"password_hash" binding:"-"`
	Money       int    `json:"money" db:"money" binding:"-"`
	IsModerator bool   `json:"-" db:"is_moderator"`
	IsFrozen    bool   `json:"-" db:"is_frozen"`
}

type UpdateUserInput struct {
	Username *string `json:"username"`
	Name     *string `json:"name"`
	Surname  *string `json:"surname"`
	Teacher  *string `json:"teacher"`
	Password *string `json:"password"`
}

type YourUser struct {
	ID          int    `json:"id" db:"id"`
	Username    string `json:"username" db:"username" binding:"required" `
	Name        string `json:"name" db:"name" binding:"required" `
	Surname     string `json:"surname" db:"surname" binding:"required"`
	Teacher     string `json:"teacher" db:"teacher" binding:"required" `
	Money       int    `json:"money" db:"money" binding:"-"`
	IsModerator bool   `json:"is_moderator" db:"is_moderator"`
	IsFrozen    bool   `json:"-" db:"is_frozen"`
}

func (input UpdateUserInput) Validate() error {
	if input.Name == nil && input.Surname == nil {
		if input.Teacher == nil && input.Password == nil {
			return errors.New("update structure has no values")
		}
	}
	return nil
}
