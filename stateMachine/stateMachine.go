package statemachine

import (
	"errors"
	"poker_server/client"
	"poker_server/command"
	"poker_server/tools"
	"sync"

	"github.com/gorilla/websocket"
)

var messageType int

var clients sync.Map

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

func DoWork(ws *websocket.Conn) {
	//读取ws中的数据
	/*
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("ReadMessage err:", err)
			break
		}

		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("WriteMessage err:", err)
			break
		}
	*/
	buffer := getMessage(ws)
	commandType := command.TypeSelector(buffer)

	if commandType == command.ADD_ROOM_REQUEST {
		if ResponseAddRoom(ws, tools.GetJsonString(buffer)) != nil {
			return
		}
	} else if commandType == command.REGISTER_ROOM_REQUEST {
		if ResponseRegisterRoom(ws) != nil {
			return
		}
	} else {
		return
	}

}

func ResponseRegisterRoom(ws *websocket.Conn) error {
	roomNum := tools.RoomNumGenerator()
	pair := [2]*client.Client{client.NewClient(), nil}
	clients.Store(roomNum, pair)
	return sendMessge(ws, string(command.RegisterRoomResponseSerialize(roomNum)))
}

func ResponseAddRoom(ws *websocket.Conn, str string) error {
	roomNum := command.AddRoomRequestDeSerialize(str)
	obj, ok := clients.Load(roomNum)
	if ok {
		pair, _ := obj.([2]*client.Client)
		pair[1] = client.NewClient()
		clients.Store(roomNum, pair)
		return sendMessge(ws, command.ResponseSerialize(command.RESULT_OK))
	} else {
		sendMessge(ws, command.ResponseSerialize(command.RESULT_NOK))
		return errors.New("no room exit")
	}
}
