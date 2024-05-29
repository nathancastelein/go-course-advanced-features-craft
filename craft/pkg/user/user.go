package user

import (
	"database/sql"
)

type User struct {
	Firstname string
	Lastname  string
}

func List(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`SELECT firstname, lastname FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Firstname, &user.Lastname); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
