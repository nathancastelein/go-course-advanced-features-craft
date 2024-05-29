package main

import (
	"database/sql"
	"log/slog"

	"github.com/nathancastelein/go-course-advanced-craft/advanced-features/solution/hexagonal/application/user"
	"github.com/nathancastelein/go-course-advanced-craft/advanced-features/solution/hexagonal/drivenadapter"
	"github.com/nathancastelein/go-course-advanced-craft/advanced-features/solution/hexagonal/driveradapter/http"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://localhost:5432/user")
	if err != nil {
		slog.Error("error while connecting to database", slog.Any("error", err))
		return
	}

	userRepository := drivenadapter.NewUserSQL(db)
	userService := user.NewService(userRepository)

	server := http.NewServer(userService)

	if err := server.Start(); err != nil {
		slog.Error("error while starting server", slog.Any("error", err))
		return
	}
}
