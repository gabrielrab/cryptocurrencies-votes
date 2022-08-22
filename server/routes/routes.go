package routes

import (
	"cryptocurrencies-votes/controllers"
	"cryptocurrencies-votes/server/socket"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func ServerRoutes(router *gin.Engine, hub *socket.Hub) *gin.Engine {
	// healthcheck
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// coins
	router.GET("/coins", controllers.ShowCoins)
	router.POST("/coin", controllers.NewCoin)

	// votes
	router.GET("/votes", controllers.ShowVotes)
	router.GET("/votes/calculate", controllers.CalculateVotes)
	router.POST("/vote/:coin/:value", func(ctx *gin.Context) {
		controllers.NewVote(ctx, hub)
	})

	//socket
	router.GET("/ws", func(ctx *gin.Context) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			return
		}

		client := socket.NewClient(ws, hub)
		hub.AddClient(client)
		go client.Watch()
	})
	return router
}
