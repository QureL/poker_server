package statemachine

import "github.com/gorilla/websocket"

func getMessage(ws *websocket.Conn) string {
	mt, message, err := ws.ReadMessage()
	if err != nil {
		return ""
	}
	messageType = mt
	return string(message)
}

func sendMessge(ws *websocket.Conn, buffer string) error {
	return ws.WriteMessage(messageType, []byte(buffer))

}
