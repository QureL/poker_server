package router

import (
	"log"
	"net/http"
	statemachine "poker_server/stateMachine"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func doWork(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(1, err)
		return
	}
	defer ws.Close()
	statemachine.DoWork(ws)
}

func Run() {
	r := gin.Default()
	r.GET("/ping", doWork)
	r.Run(":12345")
}
