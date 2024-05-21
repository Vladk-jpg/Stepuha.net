package infrastructure

import (
	"Stepuha.net/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SupplementaryPostgres struct {
	db *sqlx.DB
}

func NewSupplementaryPostgres(db *sqlx.DB) *SupplementaryPostgres {
	return &SupplementaryPostgres{db: db}
}

func (repos *SupplementaryPostgres) GetRandomGoods(userId int) ([]entities.Good, error) {
	var goods []entities.Good
	query := fmt.Sprintf("SELECT gd.id, gd.name, gd.price, gd.picture, gd.description FROM %s gd ORDER BY RANDOM() LIMIT 15", GoodsTable)
	err := repos.db.Select(&goods, query)

	return goods, err
}
