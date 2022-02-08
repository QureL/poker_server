package command

import (
	"encoding/json"
	"fmt"
	"log"
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

func commandTypeToString(commandType int) string {
	var commandStr string

	if commandType < 10 {
		commandStr = fmt.Sprintf("0%d", commandType)
	} else {
		commandStr = fmt.Sprintf("%d", commandType)
	}
	return commandStr
}

func RegisterRoomResponseSerialize(roomNun int) string {
	buffer, _ := json.Marshal(RegisterRoomResponse{RoomNum: roomNun})
	return commandTypeToString(REGISTER_ROOM_RESPONSE) + string(buffer)
}

func ResponseSerialize(res int, responseNum int) string {
	buffer, _ := json.Marshal(Response{StatusCode: res})
	return commandTypeToString(responseNum) + string(buffer)
}

func AddRoomRequestDeSerialize(str string) int {
	var req AddRoomRequest
	err := json.Unmarshal([]byte(str), &req)
	if err != nil {
		return -1
	}
	return req.RoomNum
}

func PutCardSerialize() string {
	return commandTypeToString(COMMAND_PUT_CARD)
}

func DealCardsSerialize(cs []cards.Card) string {
	dealing := Dealing{
		Cards: cs,
	}
	buffer, _ := json.Marshal(dealing)
	return commandTypeToString(DEAL_CARD_COMMAND) + string(buffer)
}

func CardsDeSerialize(str string) []cards.Card {
	log.Println("Card comming...", str)
	var dealing Dealing
	err := json.Unmarshal([]byte(str), &dealing)
	if err != nil {
		return nil
	}
	return dealing.Cards
}

func DestopCardsSeriable(cards []cards.Card) string {
	dealing := Dealing{
		Cards: cards,
	}
	buffer, _ := json.Marshal(dealing)
	return commandTypeToString(DESTOP_CARD) + string(buffer)
}

func SuccessCommandSerialize() string {
	return commandTypeToString(SUCESS_COMMAND)
}
