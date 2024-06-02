package infrastructure

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ModerPostgres struct {
	db *sqlx.DB
}

func NewModerPostgres(db *sqlx.DB) *ModerPostgres {
	return &ModerPostgres{db: db}
}

func (repos *ModerPostgres) FreezeUser(userId int) error {
	query := fmt.Sprintf("UPDATE %s SET is_frozen = TRUE WHERE gd.id = $1", UsersTable)
	_, err := repos.db.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}

func (repos *ModerPostgres) UnfreezeUser(userId int) error {
	query := fmt.Sprintf("UPDATE %s SET is_frozen = FALSE WHERE gd.id = $1", UsersTable)
	_, err := repos.db.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}
