package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// websocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(ctx *gin.Context) {
	//ctx.Writer.Header().Set("Content-Type", "text/plain")
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read err: %v", err)
			break
		}
		log.Printf("Received: %s", msg)
		if err := conn.WriteMessage(msgType, msg); err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}
}
func main() {
	router := gin.Default()
	router.GET("/ws", handleWebSocket)
	router.Run(":8000")
}
