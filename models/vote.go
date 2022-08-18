package models

import (
	"time"

	"gorm.io/gorm"
)

type Vote struct {
	Id uint `json:"id" gorm:"primaryKey unique autoIncrementIncrement"`
	Coin string `json:"coin"`
	Value int `json:"value" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}