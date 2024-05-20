package entities

type User struct {
	ID       int    `json:"-" db:"id"`
	Username string `json:"username" db:"username" binding:"required" `
	Name     string `json:"name" db:"name" binding:"required" `
	Surname  string `json:"surname" db:"surname" binding:"required"`
	Teacher  string `json:"teacher" db:"teacher" binding:"required" `
	Password string `json:"password" db:"password_hash" binding:"required"`
}

type UpdateUserInput struct {
	Username *string `json:"username"`
	Name     *string `json:"name"`
	Surname  *string `json:"surname"`
	Teacher  *string `json:"teacher"`
	Password *string `json:"password"`
}

type Cart struct {
	UserID int
	Goods  []Good
}

type favorites struct {
	UserID int
	Goods  []Good
}

type Order struct {
	ID       int
	SellerID int
	BuyerID  int
	Goods    []Good
}
