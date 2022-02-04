package command_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"poker_server/cards"
	"poker_server/command"
	"strconv"
	"sync"
	"testing"
	"time"
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

	i, err := strconv.Atoi("014"[:2])
	fmt.Println(i, err)

	fmt.Println("01234"[2:])

	a := [2]int{0, 1}
	var m sync.Map
	m.Store(1, a)

	fmt.Println(m.Load(1))

	m.Store(1, []int{34, 5})
	fmt.Println(m.Load(1))

	obj, _ := m.Load(1)
	pair := obj.([]int)
	pair[0] = 99
	fmt.Println(m.Load(1))

	fmt.Println("-----------------------")
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		fmt.Print(rand.Intn(54), " ")
	}

}
