package statemachine

import (
	"poker_server/client"
	"poker_server/command"
	"poker_server/tools"

	"github.com/gorilla/websocket"
)

var messageType int

func BuildRoom(ws *websocket.Conn) {
	buffer := getMessage(ws)
	commandType := command.TypeSelector(buffer)
	/* client instance  */
	var c *client.Client = nil
	if commandType != command.REGISTER_ROOM_REQUEST {
		return
	}
	var roomNum int
	if c = ResponseRegisterRoom(ws, &roomNum); c == nil {
		return
	}

	startChannel := storeStartChannel(roomNum)
	stopChannel := storeStopChannel(roomNum)

	<-startChannel
	for {

		buffer = <-c.ChannelOut
		sendMessge(ws, buffer)

		buffer = getMessage(ws)
		c.ChannelIn <- buffer

		flag := false
		select {
		case <-stopChannel:
			flag = true
		default:
		}
		if flag {
			break
		}
	}
}

func AddRoom(ws *websocket.Conn) {
	buffer := getMessage(ws)
	commandType := command.TypeSelector(buffer)
	/* client instance  */
	var c *client.Client = nil

	if commandType != command.ADD_ROOM_REQUEST {
		return
	}
	roomNum := 0
	if c = ResponseAddRoom(ws, tools.GetJsonString(buffer), &roomNum); c == nil {
		return
	}

	channel := getStartChannel(roomNum)
	if channel == nil {
		return
	}
	stopChannel := getStopChannel(roomNum)
	if stopChannel == nil {
		return
	}
	pair := getClients(roomNum)
	go Bussiness(pair[0], pair[1])
	channel <- struct{}{}

	for {

		buffer = <-c.ChannelOut
		sendMessge(ws, buffer)

		buffer = getMessage(ws)
		c.ChannelIn <- buffer

		flag := false
		select {
		case <-stopChannel:
			flag = true
		default:
		}
		if flag {
			break
		}
	}
}

func ResponseRegisterRoom(ws *websocket.Conn, room *int) *client.Client {
	roomNum := tools.RoomNumGenerator()
	c := client.NewClient()
	pair := [2]*client.Client{c, nil}
	clients.Store(roomNum, pair)
	if sendMessge(ws, string(command.RegisterRoomResponseSerialize(roomNum))) != nil {
		return nil
	} else {
		*room = roomNum
		return c
	}
}

func ResponseAddRoom(ws *websocket.Conn, str string, outRomm *int) *client.Client {
	roomNum := command.AddRoomRequestDeSerialize(str)
	obj, ok := clients.Load(roomNum)
	if ok {
		pair, _ := obj.([2]*client.Client)
		pair[1] = client.NewClient()
		clients.Store(roomNum, pair)
		if sendMessge(ws, command.ResponseSerialize(command.RESULT_OK)) != nil {
			return nil
		} else {
			*outRomm = roomNum
			return pair[1]
		}
	} else {
		sendMessge(ws, command.ResponseSerialize(command.RESULT_NOK))
		return nil
	}
}
