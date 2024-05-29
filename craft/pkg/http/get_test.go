package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestListUsers(t *testing.T) {
	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT firstname, lastname FROM users`).WillReturnRows(
		sqlmock.NewRows([]string{"firstname", "lastname"}).
			AddRow("SpongeBob", "SquarePants").
			AddRow("Patrick", "Star"),
	)

	server := &Server{db: db}

	// Act
	err = server.ListUsers(c)

	// Assert
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
	require.JSONEq(t, `[{"Firstname": "SpongeBob", "Lastname": "SquarePants"}, {"Firstname": "Patrick", "Lastname": "Star"}]`, rec.Body.String())
}
