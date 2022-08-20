package routes

import (
	"cryptocurrencies-votes/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerRoutes(router *gin.Engine) *gin.Engine {
	// healthcheck
	router.GET("/", func (ctx *gin.Context)  {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// coins
	router.GET("/coins", controllers.ShowCoins)
	router.GET("/coin", controllers.NewCoin)

	// votes
	router.GET("/votes", controllers.ShowVotes)
	router.GET("/votes/calculate", controllers.CalculateVotes)
	router.POST("/vote/:coin/:value", controllers.NewVote)
	return router
}