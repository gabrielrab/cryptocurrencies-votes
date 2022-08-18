package server

import (
	"cryptocurrencies-votes/server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer() Server{
	return Server{
		port: "3000",
		server: gin.Default(),
	}
}

func (s *Server) Run(){
	router := routes.ServerRoutes(s.server)
	log.Fatal(router.Run(":"+s.port))
}