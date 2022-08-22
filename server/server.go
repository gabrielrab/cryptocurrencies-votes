package server

import (
	"cryptocurrencies-votes/database"
	"cryptocurrencies-votes/server/routes"
	"cryptocurrencies-votes/server/socket"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   os.Getenv("APP_PORT"),
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	database.DatabaseConnection()
	hub := socket.NewHub()
	router := routes.ServerRoutes(s.server, hub)
	log.Fatal(router.Run(":" + s.port))
}
