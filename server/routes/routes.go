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

	// votes
	router.GET("/votes", controllers.ShowVotes)
	router.POST("/vote/:coin/:value", controllers.NewVote)
	return router
}