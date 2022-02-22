package logic

import "poker_server/cards"

func Compare(cs1, cs2 []cards.Card) int {
	if cs2 == nil {
		if ValidTest(cs1) != INVALID {
			return FIRST
		} else {
			return INCOMPARABLE
		}
	}
	if cs1 == nil {
		if ValidTest(cs2) != INVALID {
			return SECOND
		} else {
			return INCOMPARABLE
		}
	}

	cs1Type := ValidTest(cs1)
	cs2Type := ValidTest(cs2)

	if cs1Type != cs2Type {
		if cs1Type == BOMB {
			return FIRST
		} else if cs2Type == BOMB {
			return SECOND
		} else {
			return INCOMPARABLE
		}
	}
	var f func([]cards.Card, []cards.Card) int
	switch cs1Type {
	case SINGLE:
		f = SingleCompare
	case PIRE:
		f = PireCompare
	case TRIPLE:
		f = TripletCompare
	case THREE_AND_ONE:
		f = ThreeAndOneCompare
	case FOUR_AND_TWO:
		f = FourAndTwoCompare
	case BOMB:
		f = BombCompare
	case DRAGON:
		f = DragonCompare
	case PLANE:
		f = PireCompare
	case LIANDUI:
		f = LianduiCompare
	default:
		return INCOMPARABLE
	}
	return f(cs1, cs2)
}
