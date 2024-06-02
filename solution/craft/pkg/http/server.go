package http

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/nathancastelein/go-course-advanced-craft/solution/craft/pkg/user"
)

type Server struct {
	listeningPort int
	router        *echo.Echo
	user          user.Lister
}

func NewServer(listeningPort int, user user.Lister) *Server {
	router := echo.New()
	server := &Server{
		listeningPort: listeningPort,
		router:        router,
		user:          user,
	}
	router.GET("/user", server.ListUsers)
	return server
}

func (s *Server) Start() error {
	return s.router.Start(fmt.Sprintf(":%d", s.listeningPort))
}
