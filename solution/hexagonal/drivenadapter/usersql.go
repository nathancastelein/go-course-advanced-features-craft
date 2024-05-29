package drivenadapter

import (
	"context"
	"database/sql"

	"github.com/nathancastelein/go-course-advanced-craft/advanced-features/solution/hexagonal/application/user"
)

var (
	_ user.Repository = &UserSQL{}
)

type UserSQL struct {
	db *sql.DB
}

func NewUserSQL(db *sql.DB) user.Repository {
	return &UserSQL{db: db}
}

func (u *UserSQL) List(ctx context.Context) ([]user.User, error) {
	rows, err := u.db.Query(`SELECT firstname, lastname, email FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]user.User, 0)
	for rows.Next() {
		var user user.User
		if err := rows.Scan(&user.Firstname, &user.Lastname, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *UserSQL) Save(ctx context.Context, user user.User) error {
	_, err := u.db.Exec(`INSERT INTO users(firstname, lastname, email) VALUES ($1, $2, $3)`, user.Firstname, user.Lastname, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserSQL) EmailAlreadyExists(ctx context.Context, email string) (bool, error) {
	rows, err := u.db.Query(`SELECT true FROM users WHERE email = $1`, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	defer rows.Close()

	return true, nil
}
