package statemachine

import (
	"log"
	"poker_server/client"
	"poker_server/command"
	"poker_server/config"
	"poker_server/network"
	"poker_server/tools"
	"time"

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
	defer rmStartChannel(roomNum)

	timer := time.NewTimer(time.Second * config.REGISTER_ROOM_WAIT_TIME)

	select {
	case <-startChannel:
	case <-timer.C:
		log.Println("超时....")
		return
	}

	go func() {
		for {
			b := <-c.ChannelOut
			network.SendMessge(ws, b)
			select {
			case <-stopChannel:
				return
			default:
			}
		}
	}()

	go func() {
		for {
			b := network.GetMessage(ws)
			c.ChannelIn <- b
			select {
			case <-stopChannel:
				return
			default:
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

	startChannel := getStartChannel(roomNum)
	if startChannel == nil {
		return
	}
	stopChannel := getStopChannel(roomNum)
	if stopChannel == nil {
		return
	}
	pair := getClients(roomNum)
	/* start bussiness goroutine */
	go Bussiness(pair[0], pair[1], stopChannel)

	close(startChannel)

	go func() {
		for {
			b := <-c.ChannelOut
			network.SendMessge(ws, b)
			select {
			case <-stopChannel:
				return
			default:
			}
		}
	}()

	go func() {
		for {
			b := network.GetMessage(ws)
			c.ChannelIn <- b
			select {
			case <-stopChannel:
				return
			default:
			}
		}
	}()

	<-stopChannel
	rmStopChannel(roomNum)
	rmClients(roomNum)
	log.Println("end...")
}

func ResponseRegisterRoom(ws *websocket.Conn, room *int) *client.Client {
	roomNum := tools.RoomNumGenerator()
	c := client.NewClient()
	pair := [2]*client.Client{c, nil}
	clients.Store(roomNum, pair)
	time.Sleep(time.Second)
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
	time.Sleep(time.Second)
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
