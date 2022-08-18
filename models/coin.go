package models

import (
	"time"

	"gorm.io/gorm"
)

type Coin struct {
	Id uint `json:"id" gorm:"primaryKey unique autoIncrementIncrement"`
	Name string `json:"name"`
	Code string `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}