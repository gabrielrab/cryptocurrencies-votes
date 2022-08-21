package socket

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SocketConnection(router *gin.Engine) *websocket.Conn {
	var wsconnection *websocket.Conn
	router.GET("/ws", func(ctx *gin.Context) {
		ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer ws.Close()

		wsconnection = ws

		// for {
		// 	mt, message, err := ws.ReadMessage()
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		break
		// 	}

		// 	if string(message) == "ping" {
		// 		message = []byte("pong")
		// 	}

		// 	err = ws.WriteMessage(mt, message)

		// 	if err != nil {
		// 		fmt.Println(err)
		// 		break
		// 	}
		// }
	})
	return wsconnection
}
