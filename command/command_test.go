package command_test

import (
	"encoding/json"
	"fmt"
	"poker_server/cards"
	"poker_server/command"
	"testing"
)

func TestCommand(t *testing.T) {
	t.Log(command.ADD_ROOM_REQUEST)

	arr, err := json.Marshal(&command.CardRequst{
		[]cards.Card{
			{CardNum: 1},
			{CardNum: 2},
			{CardNum: 3},
		},
	})
	fmt.Println(string(arr), err)
	fmt.Println("-------------")
	str := `{"Cards":[{"CardNum":1},{"CardNum":2},{"CardNum":3}]}`
	c := new(command.CardRequst)
	err = json.Unmarshal([]byte(str), c)
	fmt.Println(err, c)
}
