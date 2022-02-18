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

	DESTOP_CARD = 10

	PASS_REQUSET  = 11
	PASS_RESPONSE = 14 /* 允许过牌 */

	SUCESS_COMMAND = 12
	FAIL_COMMAND   = 13

	SUCCESS_REQUEST = 14
)

func TypeSelector(buffer string) int {
	if len(buffer) < 2 {
		return -1
	}
	commmand_type, err := strconv.Atoi(buffer[:2])
	if err != nil {
		return -1
	} else {
		return commmand_type
	}
}
