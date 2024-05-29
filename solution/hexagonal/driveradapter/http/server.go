package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nathancastelein/go-course-advanced-craft/advanced-features/solution/hexagonal/application/user"
)

type Server struct {
	userService user.Service
}

func NewServer(userService user.Service) *Server {
	return &Server{
		userService: userService,
	}
}

func (s *Server) Start() error {
	router := echo.New()
	router.GET("/user", s.listUsers)
	router.POST("/user", s.createUser)

	return router.Start(":8080")
}

func (s *Server) listUsers(ctx echo.Context) error {
	users, err := s.userService.List(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newBodyError(err))
	}

	return ctx.JSON(http.StatusOK, users)
}

type createUserBody struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

func (s *Server) createUser(ctx echo.Context) error {
	body := createUserBody{}
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, newBodyError(err))
	}

	user, err := s.userService.Create(ctx.Request().Context(), body.Firstname, body.Lastname, body.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newBodyError(err))
	}

	return ctx.JSON(http.StatusCreated, user)
}
