package network

import "github.com/gorilla/websocket"

var messageType int

func GetMessage(ws *websocket.Conn) string {
	mt, message, err := ws.ReadMessage()
	if err != nil {
		return ""
	}
	messageType = mt
	return string(message)
}

func SendMessge(ws *websocket.Conn, buffer string) error {
	return ws.WriteMessage(messageType, []byte(buffer))

}
