package tools_test

import (
	"fmt"
	"poker_server/cards"
	"poker_server/tools"
	"sort"
	"sync"
	"testing"
)

func TestTools(t *testing.T) {
	var m sync.Map
	c := &cards.Card{CardNum: 90}
	m.Store(1, c)
	ret, _ := m.Load(1)

	card, _ := ret.(*cards.Card)

	fmt.Println(card.CardNum)

	m.Delete(1)

	fmt.Println(m.Load(9))

	fmt.Println("---------------------------")

	var arr []int
	for i := 1; i <= 54; i++ {
		arr = append(arr, i)
	}
	arr = tools.Shuffle(arr)
	arr = arr[10:]
	fmt.Println(arr)
	fmt.Println("---------------------------")

	j := arr[:44/2]
	fmt.Println(len(j), j)

	k := arr[44/2:]
	fmt.Println(len(k), k)
	fmt.Println("---------------------------")

	sort.Ints(j)
	sort.Ints(k)
	fmt.Println(len(j), j)
	fmt.Println(len(k), k)
	fmt.Println(arr)

}
