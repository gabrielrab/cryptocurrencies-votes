package server

import (
	"cryptocurrencies-votes/database"
	"cryptocurrencies-votes/server/routes"
	"cryptocurrencies-votes/server/socket"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "3000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	database.DatabaseConnection()
	websocket := socket.SocketConnection(s.server)
	router := routes.ServerRoutes(s.server, websocket)
	log.Fatal(router.Run(":" + s.port))
}
