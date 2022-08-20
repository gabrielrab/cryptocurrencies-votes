package controllers

import (
	"cryptocurrencies-votes/database"
	"cryptocurrencies-votes/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowCoins(ctx *gin.Context){
	db := database.GetDatabase()
	var coins []models.Coin

	err := db.Find(&coins).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Error when show coins",
		})
		return
	}

	ctx.JSON(http.StatusOK, coins)
}

func NewCoin (ctx *gin.Context){
	db := database.GetDatabase()
	var coin models.Coin

	err := ctx.ShouldBind(&coin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Error when create coin",
		})
		return
	}

	err = db.Create(&coin).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Error when create coin: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, coin)
}