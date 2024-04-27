package infrastructure

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	UsersTable      = "users"
	GoodsTable      = "goods"
	UsersGoodsTable = "users_goods"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	DBName   string
	Password string
	SSLMode  string
}

func NewPostrgesDB(conf DbConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		conf.Host, conf.Port, conf.Username, conf.DBName, conf.Password, conf.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
