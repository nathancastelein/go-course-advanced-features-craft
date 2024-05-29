package http

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/nathancastelein/go-course-advanced-craft/advanced-features/solution/hexagonal/application/user"
	"github.com/stretchr/testify/require"
)

var (
	_ user.Service = &UserServiceStub{}
)

type UserServiceStub struct{}

func (u *UserServiceStub) List(ctx context.Context) ([]user.User, error) {
	return []user.User{
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
	}, nil
}

func (u *UserServiceStub) Create(ctx context.Context, firstname, lastname, email string) (user.User, error) {
	return user.User{
		Firstname: "SpongeBob",
		Lastname:  "SquarePants",
		Email:     "sponge@bob.com",
	}, nil
}

func TestListUsers(t *testing.T) {
	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	server := NewServer(&UserServiceStub{})

	// Act
	err := server.listUsers(c)

	// Assert
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
	require.JSONEq(t, `[{"Firstname": "SpongeBob", "Lastname": "SquarePants", "Email": "sponge@bob.com"}, {"Firstname": "Patrick", "Lastname": "Star", "Email": "patrick@star.com"}]`, rec.Body.String())
}

func TestCreateUser(t *testing.T) {
	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBufferString(`{"firstname": "SpongeBob", "lastname": "SquarePants", "email": "sponge@bob.com"}`))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	server := NewServer(&UserServiceStub{})

	// Act
	err := server.createUser(c)

	// Assert
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, rec.Code)
	require.JSONEq(t, `{"Firstname": "SpongeBob", "Lastname": "SquarePants", "Email": "sponge@bob.com"}`, rec.Body.String())
}
