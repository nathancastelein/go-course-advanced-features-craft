package user

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestListUsers(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT firstname, lastname FROM users`).WillReturnRows(
		sqlmock.NewRows([]string{"firstname", "lastname"}).
			AddRow("SpongeBob", "SquarePants").
			AddRow("Patrick", "Star"),
	)

	// Act
	users, err := List(db)

	// Assert
	require.NoError(t, err)
	require.Equal(t, []User{
		{
			Firstname: "SpongeBob",
			Lastname:  "SquarePants",
		},
		{
			Firstname: "Patrick",
			Lastname:  "Star",
		},
	}, users)
}
