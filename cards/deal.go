package cards

import (
	"poker_server/tools"
	"sort"
)

/* 发牌逻辑 */

const (
	DECK_SIZE   = 54
	DISCARD_NUM = 10
)

func GenerateCard() ([]Card, []Card) {
	var arr []int
	for i := 1; i <= DECK_SIZE; i++ {
		arr = append(arr, i)
	}
	arr = tools.Shuffle(arr)
	arr = arr[DISCARD_NUM:]

	arr1 := arr[:(DECK_SIZE-DISCARD_NUM)/2]
	arr2 := arr[(DECK_SIZE-DISCARD_NUM)/2:]

	sort.Ints(arr1)
	sort.Ints(arr2)
	var cards1, cards2 []Card

	for _, i := range arr1 {
		cards1 = append(cards1, NewCardFromInt(i))
	}

	for _, i := range arr2 {
		cards2 = append(cards2, NewCardFromInt(i))
	}
	return cards1, cards2
}
