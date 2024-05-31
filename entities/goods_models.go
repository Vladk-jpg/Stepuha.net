package entities

import "errors"

type Good struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Price       int    `json:"price" db:"price" binding:"required"`
	Picture     string `json:"picture" db:"picture"`
	Description string `json:"description" db:"description" binding:"required"`
}

type UpdateGoodInput struct {
	Name        *string `json:"name"`
	Price       *int    `json:"price"`
	Picture     *string `json:"picture"`
	Description *string `json:"description"`
}

func (input UpdateGoodInput) Validate() error {
	if input.Name == nil && input.Picture == nil {
		if input.Description == nil && input.Price == nil {
			return errors.New("update structure has no values")
		}
	}
	return nil
}
