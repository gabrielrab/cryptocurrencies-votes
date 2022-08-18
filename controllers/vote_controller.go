package controllers

import (
	"cryptocurrencies-votes/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ShowVotes(ctx *gin.Context) {
	var votes = []models.Vote{
		{Id: 1, Coin: "BTC", Value: -1},
		{Id: 2, Coin: "BTC", Value: -1},
		{Id: 3, Coin: "USDT", Value: 1},
		{Id: 4, Coin: "KLV", Value: 1},
		{Id: 5, Coin: "KLV", Value: -1},
		{Id: 6, Coin: "ETH", Value: 1},
		{Id: 7, Coin: "USDT", Value: 1},
		{Id: 8, Coin: "BTC", Value: -1},
		{Id: 9, Coin: "ETH", Value: -1},
		{Id: 10, Coin: "KLV", Value: 1},
		{Id: 11, Coin: "ETH", Value: 1},
		{Id: 12, Coin: "USDT", Value: 1},
	}

	ctx.JSON(http.StatusOK, votes)
}

func NewVote(ctx *gin.Context){
	voteValue := ctx.Param("value")
	if(voteValue == "" || (voteValue != "up" && voteValue != "down")){
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid vote value",
		})
		return
	}

	newVote := models.Vote{
		Coin: strings.ToUpper(ctx.Param("coin")),
		Value: func () int {
			if (voteValue == "up"){
				return 1
			} else{
				return -1
			}
		}(),
	}
	ctx.JSON(http.StatusOK, newVote)
}