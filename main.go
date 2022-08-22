package main

import (
	"cryptocurrencies-votes/server"
)

func main() {
	s := server.NewServer()
	s.Run()
}
