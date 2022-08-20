package controllers

import (
	"cryptocurrencies-votes/database"
	"cryptocurrencies-votes/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ShowVotes(ctx *gin.Context) {
	db := database.GetDatabase()
	var votes []models.Vote

	err := db.Find(&votes).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Error when show votes",
		})
	}

	ctx.JSON(http.StatusOK, votes)
}

func CalculateVotes(ctx *gin.Context){
	var calculatedVotes []struct{
		Code string `json:"code"`
		Name string `json:"name"`
		Votes int `json:"votes"`
	}

	db := database.GetDatabase()
	err := db.Raw(`
			select  
				coins.code, 
				coins.name,
				sum(votes.value) as votes
			from votes
				inner join coins on votes.coin = coins.code
			group by coins.code, coins.name`).Scan(&calculatedVotes).Error
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Error on calculate votes",
		})
		return
	}

	ctx.JSON(http.StatusOK, calculatedVotes)
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

	db := database.GetDatabase()
	var coin models.Coin

	err := db.Where(&models.Coin{ Code: strings.ToUpper(ctx.Param("coin")) }).First(&coin).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": "Coin not found.",
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

	
	err = db.Create(&newVote).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Error when create vote",
		})
		return
	}

	ctx.JSON(http.StatusCreated, newVote)
}