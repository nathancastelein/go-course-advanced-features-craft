package adapters

import (
	"database/sql"

	"github.com/nathancastelein/go-course-advanced-craft/solution/craft/pkg/user"
)

type UserSQL struct {
	db *sql.DB
}

func NewUserSQL(db *sql.DB) user.Lister {
	return &UserSQL{
		db: db,
	}
}

func (u *UserSQL) List() ([]user.User, error) {
	rows, err := u.db.Query(`SELECT firstname, lastname FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]user.User, 0)
	for rows.Next() {
		var user user.User
		if err := rows.Scan(&user.Firstname, &user.Lastname); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
