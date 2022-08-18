package models

type Vote struct {
	Id string `json:"id"`
	Coin string `json:"coin"`
	Value int `json:"value" binding:"required"`
}