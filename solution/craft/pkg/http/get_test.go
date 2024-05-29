package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/nathancastelein/go-course-advanced-craft/advanced-features/solution/craft/pkg/user"
	"github.com/stretchr/testify/require"
)

var (
	_ user.Lister = &UserListerStub{}
)

type UserListerStub struct{}

func (u *UserListerStub) List() ([]user.User, error) {
	return []user.User{
		{
			Firstname: "SpongeBob",
			Lastname:  "SquarePants",
		},
		{
			Firstname: "Patrick",
			Lastname:  "Star",
		},
	}, nil
}

func TestListUsers(t *testing.T) {
	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	server := &Server{user: &UserListerStub{}}

	// Act
	err := server.ListUsers(c)

	// Assert
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
	require.JSONEq(t, `[{"firstname": "SpongeBob", "lastname": "SquarePants"}, {"firstname": "Patrick", "lastname": "Star"}]`, rec.Body.String())
}
