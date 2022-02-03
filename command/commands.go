package command

import (
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
