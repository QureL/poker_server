package command

import "strconv"

const (
	REGISTER_ROOM_REQUEST  = 0
	REGISTER_ROOM_RESPONSE = 1

	ADD_ROOM_REQUEST  = 2
	ADD_ROOM_RESPONSE = 3

	CARD_REQUEST  = 4
	CARD_RESPONSE = 5

	SUCCESS_RESPONSE = 6
	FAIL_RESPONSE    = 7

	COMMAND_PUT_CARD  = 8 /* 出牌指令 */
	DEAL_CARD_COMMAND = 9 /* 服务端发牌 */
)

func TypeSelector(buffer string) int {
	commmand_type, err := strconv.Atoi(buffer[:2])
	if err != nil {
		return -1
	} else {
		return commmand_type
	}
}
