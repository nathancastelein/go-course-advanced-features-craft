package drivenadapter

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nathancastelein/go-course-advanced-craft/solution/hexagonal/application/user"
	"github.com/stretchr/testify/require"
)

func TestListUsersSQL(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT firstname, lastname, email FROM users`).WillReturnRows(
		sqlmock.NewRows([]string{"firstname", "lastname", "email"}).
			AddRow("SpongeBob", "SquarePants", "sponge@bob.com").
			AddRow("Patrick", "Star", "patrick@star.com"),
	)

	userRepository := &UserSQL{db: db}

	// Act
	users, err := userRepository.List(context.Background())

	// Assert
	require.NoError(t, err)
	require.Equal(t, []user.User{
		{
			Firstname: "SpongeBob",
			Lastname:  "SquarePants",
			Email:     "sponge@bob.com",
		},
		{
			Firstname: "Patrick",
			Lastname:  "Star",
			Email:     "patrick@star.com",
		},
	}, users)
}

func TestSaveUsersSQL(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO users(firstname, lastname, email) VALUES ($1, $2, $3)`)).
		WithArgs("SpongeBob", "SquarePants", "sponge@bob.com").
		WillReturnResult(driver.RowsAffected(1))

	userRepository := &UserSQL{db: db}

	// Act
	err = userRepository.Save(context.Background(), user.User{
		Firstname: "SpongeBob",
		Lastname:  "SquarePants",
		Email:     "sponge@bob.com",
	})

	// Assert
	require.NoError(t, err)
}

func TestEmailAlreadyExistsUsersSQL(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT true FROM users WHERE email = $1`)).
		WithArgs("sponge@bob.com").
		WillReturnRows(
			sqlmock.NewRows([]string{"?"}).
				AddRow(true),
		)

	userRepository := &UserSQL{db: db}

	// Act
	emailAlreadyExists, err := userRepository.EmailAlreadyExists(context.Background(), "sponge@bob.com")

	// Assert
	require.NoError(t, err)
	require.True(t, emailAlreadyExists)
}

func TestEmailNotAlreadyExistsUsersSQL(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT true FROM users WHERE email = $1`)).
		WithArgs("sponge@bob.com").
		WillReturnError(sql.ErrNoRows)

	userRepository := &UserSQL{db: db}

	// Act
	emailAlreadyExists, err := userRepository.EmailAlreadyExists(context.Background(), "sponge@bob.com")

	// Assert
	require.NoError(t, err)
	require.False(t, emailAlreadyExists)
}
