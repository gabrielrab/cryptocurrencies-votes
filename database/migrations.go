package database

import (
	"cryptocurrencies-votes/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB){
	db.AutoMigrate(models.Vote{})
	db.AutoMigrate(models.Coin{})
}