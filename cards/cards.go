package cards

const (
	Card_3           = 1
	Card_4           = 2
	Card_5           = 3
	Card_6           = 4
	Card_7           = 5
	Card_8           = 6
	Card_9           = 7
	Card_10          = 8
	Card_j           = 9
	Card_q           = 10
	Card_k           = 11
	Card_a           = 12
	Card_2           = 13
	Card_joker_small = 14
	Card_joker_big   = 15
)

const (
	Heart    = 1 /* 红桃 */
	Club     = 2 /* 梅花 */
	Dianmond = 3 /* 方块 */
	Spade    = 4 /* 黑桃 */
)

type Card struct {
	CardNum int `json:"CardNum"`
	Decor   int `json:"Decor"`
}

func NewCardFromInt(num int) Card {
	/*
		num = cardnum * decor
	*/
	if num == 53 {
		return Card{
			CardNum: Card_joker_small,
		}
	}
	if num == 54 {
		return Card{
			CardNum: Card_joker_big,
		}
	}
	var cardnum int
	var decor int
	if num%4 == 0 {
		cardnum = num / 4
	} else {
		cardnum = num/4 + 1
	}

	if num%4 == 0 {
		decor = 4
	} else {
		decor = num % 4
	}
	return Card{
		CardNum: cardnum,
		Decor:   decor,
	}

}
