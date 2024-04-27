package entities

type User struct {
	ID       int    `json:"-" db:"id"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Teacher  string `json:"teacher" binding:"required"`
	//Email    string
	Password string `json:"password" binding:"required"`
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
