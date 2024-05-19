package entities

type User struct {
	ID       int    `json:"-" db:"id"`
	Username string `json:"username" db:"username" binding:"required" `
	Name     string `json:"name" db:"name" binding:"required" `
	Surname  string `json:"surname" db:"surname" binding:"required"`
	Teacher  string `json:"teacher" db:"teacher" binding:"required" `
	//Email    string
	Password string `json:"password" db:"password_hash" binding:"required"`
	//Rating   int
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
