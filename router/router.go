package router

import (
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

func buildroom(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	statemachine.BuildRoomHandler(ws)
}

func addroom(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	statemachine.AddRoomHandler(ws)
}

func Run() {
	r := gin.Default()
	r.GET("/buildroom", buildroom)
	r.GET("/addroom", addroom)
	r.Run(":12346")
}
