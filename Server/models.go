package main

type User struct {
	ID       int
	Name     string
	Surname  string
	Email    string
	Password string
	Teacher  string
	//Rating   int
}

type Good struct {
	ID          int
	Name        string
	Price       float64
	Picture     string
	Category    string
	Description string
}

type Cart struct {
	UserID int
	Goods  []Good
}

type favotites struct {
	UserID int
	Goods  []Good
}

type Order struct {
	ID       int
	SellerID int
	BuyerID  int
	Goods    []Good
}
