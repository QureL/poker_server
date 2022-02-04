package command

import (
	"encoding/json"
	"poker_server/cards"
)

const (
	RESULT_OK  = 0
	RESULT_NOK = 1
)

type Response struct {
	StatusCode int `json:"StatusCode"`
}

type AddRoomRequest struct {
	RoomNum int `json:"RoomNum"`
}

type RegisterRoomResponse struct {
	RoomNum int `json:"RoomNum"`
}

type CardRequst struct {
	Cards []cards.Card `json:"Cards"`
}

/* 发牌 */
type Dealing struct {
	Cards []cards.Card `json:"Cards"`
}

func RegisterRoomResponseSerialize(roomNun int) string {
	buffer, _ := json.Marshal(RegisterRoomResponse{RoomNum: roomNun})
	return string(buffer)
}

func ResponseSerialize(res int) string {
	buffer, _ := json.Marshal(Response{StatusCode: res})
	return string(buffer)
}

func AddRoomRequestDeSerialize(str string) int {
	var req AddRoomRequest
	err := json.Unmarshal([]byte(str), &req)
	if err != nil {
		return -1
	}
	return req.RoomNum
}
