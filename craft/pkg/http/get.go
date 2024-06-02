package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nathancastelein/go-course-advanced-craft/craft/pkg/user"
)

func (s *Server) ListUsers(ctx echo.Context) error {
	users, err := user.List(s.db)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, users)
}
