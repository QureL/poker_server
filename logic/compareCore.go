package logic

import (
	"poker_server/cards"
	"sort"
)

const (
	FIRST        = 1
	SECOND       = 2
	INCOMPARABLE = -1
)

func SingleCompare(cs1, cs2 []cards.Card) int {
	if cs1[0].CardNum == cs2[0].CardNum {
		return INCOMPARABLE
	} else if cs1[0].CardNum > cs2[0].CardNum {
		return FIRST
	} else {
		return SECOND
	}
}

func PireCompare(cs1, cs2 []cards.Card) int {
	if cs1[0].CardNum == cs2[0].CardNum {
		return INCOMPARABLE
	} else if cs1[0].CardNum > cs2[0].CardNum {
		return FIRST
	} else {
		return SECOND
	}
}

func BombCompare(cs1, cs2 []cards.Card) int {
	if cs1[0].CardNum == cs2[0].CardNum {
		return INCOMPARABLE
	} else if cs1[0].CardNum > cs2[0].CardNum {
		return FIRST
	} else {
		return SECOND
	}
}

func TripletCompare(cs1, cs2 []cards.Card) int {
	if cs1[0].CardNum == cs2[0].CardNum {
		return INCOMPARABLE
	} else if cs1[0].CardNum > cs2[0].CardNum {
		return FIRST
	} else {
		return SECOND
	}
}

func ThreeAndOneCompare(cs1, cs2 []cards.Card) int {
	var cs1IntArr, cs2IntArr []int

	for _, value := range cs1 {
		cs1IntArr = append(cs1IntArr, value.CardNum)
	}
	for _, value := range cs2 {
		cs2IntArr = append(cs2IntArr, value.CardNum)
	}

	sort.Ints(cs1IntArr)
	sort.Ints(cs2IntArr)

	var cs1Flag, cs2Flag int
	if cs1IntArr[0] == cs1IntArr[1] && cs1IntArr[1] == cs1IntArr[2] && cs1IntArr[1] != cs1IntArr[3] {
		cs1Flag = cs1IntArr[0]
	} else {
		cs1Flag = cs1IntArr[3]
	}

	if cs2IntArr[0] == cs2IntArr[1] && cs2IntArr[1] == cs2IntArr[2] && cs2IntArr[1] != cs2IntArr[3] {
		cs2Flag = cs2IntArr[0]
	} else {
		cs2Flag = cs2IntArr[3]
	}

	if cs1Flag > cs2Flag {
		return FIRST
	} else {
		return SECOND
	}
}

func FourAndTwoCompare(cs1, cs2 []cards.Card) int {
	var cs1IntArr, cs2IntArr []int

	for _, value := range cs1 {
		cs1IntArr = append(cs1IntArr, value.CardNum)
	}
	for _, value := range cs2 {
		cs2IntArr = append(cs2IntArr, value.CardNum)
	}

	sort.Ints(cs1IntArr)
	sort.Ints(cs2IntArr)

	var cs1Flag, cs2Flag int

	if cs1IntArr[0] == cs1IntArr[1] && cs1IntArr[1] == cs1IntArr[2] && cs1IntArr[2] == cs1IntArr[3] && cs1IntArr[3] != cs1IntArr[4] && cs1IntArr[4] == cs1IntArr[5] {
		cs1Flag = cs1IntArr[0]
	} else {
		cs1Flag = cs1IntArr[5]
	}

	if cs2IntArr[0] == cs2IntArr[1] && cs2IntArr[1] == cs2IntArr[2] && cs2IntArr[2] == cs2IntArr[3] && cs2IntArr[3] != cs2IntArr[4] && cs2IntArr[4] == cs2IntArr[5] {
		cs2Flag = cs2IntArr[0]
	} else {
		cs2Flag = cs2IntArr[5]
	}
	if cs1Flag > cs2Flag {
		return FIRST
	} else {
		return SECOND
	}
}

func DragonCompare(cs1, cs2 []cards.Card) int {
	if len(cs1) != len(cs2) {
		return INCOMPARABLE
	}
	var cs1IntArr, cs2IntArr []int
	for _, value := range cs1 {
		cs1IntArr = append(cs1IntArr, value.CardNum)
	}
	for _, value := range cs2 {
		cs2IntArr = append(cs2IntArr, value.CardNum)
	}

	sort.Ints(cs1IntArr)
	sort.Ints(cs2IntArr)
	if cs1IntArr[0] > cs2IntArr[0] {
		return FIRST
	} else if cs1IntArr[0] < cs2IntArr[0] {
		return SECOND
	} else {
		return INCOMPARABLE
	}
}
