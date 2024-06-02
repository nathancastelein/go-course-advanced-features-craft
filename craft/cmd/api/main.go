package main

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/nathancastelein/go-course-advanced-craft/craft/pkg/http"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://localhost:5432/user")
	if err != nil {
		slog.Error("error while opening database", slog.Any("error", err))
		return
	}

	server := http.NewServer(8080, db)
	if err := server.Start(); err != nil {
		slog.Error("error while starting server", slog.Any("error", err))
		return
	}
}
