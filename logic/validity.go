package logic

import (
	"poker_server/cards"
	"sort"
	"testing"
)

const (
	FIRST  = 1
	SECOND = 2
)

const ( /* 牌组类型 */
	SINGLE        = 1
	PIRE          = 2
	TRIPLE        = 3
	BOMB          = 4
	THREE_AND_ONE = 5
	DRAGON        = 6
	FOUR_AND_TWO  = 7
	INVALID       = -1
)

func ValidTest(cards1 []cards.Card, cards2 []cards.Card) int {

	return 0
}

func isBomb(cs []cards.Card) bool {
	if len(cs) == 4 {
		tmp := cs[0]
		for _, i := range cs {
			if i != tmp {
				return false
			}
		}
		return true
	}
	return false
}

func isSingle(cs []cards.Card) bool {
	if len(cs) == 1 {
		return true
	}
	return false
}

func isPire(cs []cards.Card) bool {
	if len(cs) == 2 {
		if cs[0] == cs[1] {
			return true
		}
	}
	return false
}

func isTriplet(cs []cards.Card) bool {
	if len(cs) != 3 {
		return false
	}
	tmp := cs[0]
	for _, i := range cs {
		if i != tmp {
			return false
		}
	}
	return true
}

func isThreeAndOne(cs []cards.Card) bool {
	if len(cs) != 4 {
		return false
	}
	var arr []int
	for _, v := range cs {
		arr = append(arr, v.CardNum)
	}
	sort.Ints(arr)
	if arr[0] == arr[1] && arr[1] == arr[2] && arr[1] != arr[3] {
		return true
	}
	if arr[0] != arr[1] && arr[1] == arr[2] && arr[2] == arr[3] {
		return true
	}
	return false
}

func isDragon(cs []cards.Card) bool {
	if len(cs) < 5 {
		return false
	}

	var arrInt []int

	for _, value := range cs {
		arrInt = append(arrInt, value.CardNum)
	}

	sort.Ints(arrInt)

	if arrInt[-1] 
}
