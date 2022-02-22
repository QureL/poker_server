package logic

import (
	"poker_server/cards"
	"sort"
)

const ( /* 牌组类型 */
	SINGLE        = 1
	PIRE          = 2
	TRIPLE        = 3
	BOMB          = 4
	THREE_AND_ONE = 5
	DRAGON        = 6
	FOUR_AND_TWO  = 7
	PLANE         = 8
	LIANDUI       = 9
	INVALID       = -1
)

func getCardsNums(cs []cards.Card) []int {
	var arr []int = make([]int, 0, len(cs))
	for _, v := range cs {
		arr = append(arr, v.CardNum)
	}
	sort.Ints(arr)
	return arr
}

func ValidTest(cs []cards.Card) int {
	cards_length := len(cs)
	switch {
	case cards_length == 1:
		if isSingle(cs) {
			return SINGLE
		} else {
			return INVALID
		}
	case cards_length == 2:
		if isPire(cs) {
			return PIRE
		} else {
			return INVALID
		}

	case cards_length == 3:
		if isTriplet(cs) {
			return TRIPLE
		} else {
			return INVALID
		}
	case cards_length == 4:
		if isBomb(cs) {
			return BOMB
		} else if isThreeAndOne(cs) {
			return THREE_AND_ONE
		} else {
			return INVALID
		}

	case cards_length >= 5:
		if isFourAndTwo(cs) {
			return FOUR_AND_TWO
		} else if isDragon(cs) {
			return DRAGON
		} else if isPlane(cs, nil) {
			return PLANE
		} else if isLiandui(cs, nil) {
			return LIANDUI
		} else {
			return INVALID
		}
	}
	return INVALID
}

func isBomb(cs []cards.Card) bool {
	if len(cs) == 4 {
		arr := getCardsNums(cs)
		tmp := arr[0]

		for _, i := range arr {
			if i != tmp {
				return false
			}
		}
		return true
	}
	return false
}

func isSingle(cs []cards.Card) bool {
	return len(cs) == 1
}

func isPire(cs []cards.Card) bool {
	if len(cs) != 2 {
		return false
	}
	arr := getCardsNums(cs)
	if arr[0] == arr[1] {
		return true
	} else {
		return false
	}
}

func isTriplet(cs []cards.Card) bool {
	if len(cs) != 3 {
		return false
	}
	arr := getCardsNums(cs)
	tmp := arr[0]
	for _, i := range arr {
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
	var arr []int = getCardsNums(cs)
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
	var arrInt []int = getCardsNums(cs)
	if arrInt[len(arrInt)-1] >= cards.Card_a {
		return false
	}

	index := 1
	for ; index < len(arrInt); index++ {
		if arrInt[index]-arrInt[index-1] != 1 {
			return false
		}
	}
	return true
}

func isFourAndTwo(cs []cards.Card) bool {
	if len(cs) != 6 {
		return false
	}
	var arr []int
	for _, v := range cs {
		arr = append(arr, v.CardNum)
	}
	sort.Ints(arr)
	if arr[0] == arr[1] && arr[1] == arr[2] && arr[2] == arr[3] && arr[3] != arr[4] && arr[4] == arr[5] {
		return true
	}
	if arr[0] == arr[1] && arr[1] != arr[2] && arr[2] == arr[3] && arr[3] == arr[4] && arr[4] == arr[5] {
		return true
	}
	return false
}

func isPlane(cs []cards.Card, bigest *int) bool {
	if len(cs)%4 != 0 && len(cs) < 8 {
		return false
	}
	arr := getCardsNums(cs)

	m_cards := make(map[int]int)

	for _, v := range arr {
		m_cards[v] += 1
	}
	cnt := 0

	var card_three []int
	for i, v := range m_cards {
		if v >= 3 {
			card_three = append(card_three, i)
			cnt++
		}
	}
	if cnt != len(cs)/4 {
		return false
	}

	sort.Ints(card_three)

	for i := 1; i < len(card_three); i++ {
		if card_three[i]-card_three[i-1] != 1 {
			return false
		}
	}

	if card_three[len(card_three)-1] >= cards.Card_2 {
		return false
	}

	if bigest != nil {
		*bigest = card_three[len(card_three)-1]
	}
	return true
}

func isLiandui(cs []cards.Card, bigest *int) bool {
	if len(cs) < 6 || len(cs)%2 != 0 {
		return false
	}
	arr := getCardsNums(cs)

	m_cards := make(map[int]int)

	for _, v := range arr {
		m_cards[v] += 1
	}

	var tmp []int

	for i, v := range m_cards {
		if v != 2 {
			return false
		}
		tmp = append(tmp, i)
	}

	sort.Ints(tmp)

	for i := 1; i < len(tmp); i++ {
		if tmp[i]-tmp[i-1] != 1 {
			return false
		}
	}

	if bigest != nil {
		*bigest = tmp[len(tmp)-1]
	}
	return true
}
