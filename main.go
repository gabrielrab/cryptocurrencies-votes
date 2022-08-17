package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Coin struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Vote struct {
	Id string `json:"id"`
	Coin string `json:"coin"`
	Value int `json:"value" binding:"required"`
}

// populate
var votes = []Vote{
	{Id: "1", Coin: "BTC", Value: -1},
	{Id: "2", Coin: "BTC", Value: -1},
	{Id: "3", Coin: "USDT", Value: 1},
	{Id: "4", Coin: "KLV", Value: 1},
	{Id: "5", Coin: "KLV", Value: -1},
	{Id: "6", Coin: "ETH", Value: 1},
	{Id: "7", Coin: "USDT", Value: 1},
	{Id: "8", Coin: "BTC", Value: -1},
	{Id: "9", Coin: "ETH", Value: -1},
	{Id: "10", Coin: "KLV", Value: 1},
	{Id: "11", Coin: "ETH", Value: 1},
	{Id: "12", Coin: "USDT", Value: 1},
}

func main(){
	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	server.GET("/vote", getVotes)
	server.POST("/vote/:coin/:value", newVote)


	server.Run("localhost:3000")
}

func getVotes(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, votes)
}

func newVote(ctx *gin.Context){
	voteValue := ctx.Param("value")

	if(voteValue == "" || (voteValue != "up" && voteValue != "down")){
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid vote value",
		})
		return
	}

	newVote := Vote{
		Coin: ctx.Param("coin"),
		Value: func () int {
			if (voteValue == "up"){
				return 1
			} else{
				return -1
			}
		}(),
	}

	fmt.Println(newVote)
	ctx.IndentedJSON(http.StatusCreated, &newVote)

}