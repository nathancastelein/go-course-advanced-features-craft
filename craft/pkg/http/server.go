package http

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	listeningPort int
	router        *echo.Echo
	db            *sql.DB
}

func NewServer(listeningPort int, db *sql.DB) *Server {
	router := echo.New()
	server := &Server{
		listeningPort: listeningPort,
		router:        router,
		db:            db,
	}
	router.GET("/user", server.ListUsers)
	return server
}

func (s *Server) Start() error {
	return s.router.Start(fmt.Sprintf(":%d", s.listeningPort))
}
