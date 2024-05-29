package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) ListUsers(ctx echo.Context) error {
	users, err := s.user.List()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	jsonUsers := make([]User, len(users))
	for i, user := range users {
		jsonUsers[i] = (&User{}).From(user)
	}

	return ctx.JSON(http.StatusOK, jsonUsers)
}
