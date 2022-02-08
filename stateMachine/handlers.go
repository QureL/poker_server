package statemachine

import (
	"log"
	"poker_server/client"
	"poker_server/command"
	"poker_server/network"
	"poker_server/tools"

	"github.com/gorilla/websocket"
)

func BuildRoomHandler(ws *websocket.Conn) {
	buffer := network.GetMessage(ws)
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

	go func() {
		for {
			b := <-c.ChannelOut
			network.SendMessge(ws, b)

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
	}()

	go func() {
		for {
			b := network.GetMessage(ws)
			c.ChannelIn <- b

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
	}()

	<-stopChannel
	log.Println("end...")
}

func AddRoomHandler(ws *websocket.Conn) {
	buffer := network.GetMessage(ws)
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
	/* start bussiness goroutine */
	go Bussiness(pair[0], pair[1])

	channel <- struct{}{}

	go func() {
		for {
			b := <-c.ChannelOut
			network.SendMessge(ws, b)

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
	}()

	go func() {
		for {
			b := network.GetMessage(ws)
			c.ChannelIn <- b

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
	}()

	<-stopChannel
	log.Println("end...")
}

func ResponseRegisterRoom(ws *websocket.Conn, room *int) *client.Client {
	roomNum := tools.RoomNumGenerator()
	c := client.NewClient()
	pair := [2]*client.Client{c, nil}
	clients.Store(roomNum, pair)
	if network.SendMessge(ws, string(command.RegisterRoomResponseSerialize(roomNum))) != nil {
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
		if network.SendMessge(ws, command.ResponseSerialize(command.RESULT_OK, command.ADD_ROOM_RESPONSE)) != nil {
			return nil
		} else {
			*outRomm = roomNum
			return pair[1]
		}
	} else {
		network.SendMessge(ws, command.ResponseSerialize(command.RESULT_NOK, command.ADD_ROOM_RESPONSE))
		return nil
	}
}
