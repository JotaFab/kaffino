package server

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"kaffino/internal/database"
)

type Server struct {
	port int

	db database.Service
}

func NewServer() *http.Server {
	port := 8080 
	NewServer := &Server{
		port: port,

		db: database.NewDB(),
	}
	err := NewServer.db.DbInit()
	if err != nil {
		fmt.Println(err)
	}
	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
