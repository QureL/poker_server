package statemachine

import (
	"log"
	"math/rand"
	"poker_server/cards"
	"poker_server/client"
	"poker_server/command"
	"poker_server/logic"
	"time"
)

func Bussiness(c1, c2 *client.Client) {
	cards1, cards2 := cards.GenerateCard()

	var destop, newCards []cards.Card
	var buffer string
	var tmp *client.Client
	/* 给客户端建立连接的时间 */
	time.Sleep(time.Second * 2)
	dealCards(cards1, c1)
	dealCards(cards2, c2)
	fisrt, other := disideCardFirst(c1, c2)

start:
	buffer = <-fisrt.ChannelIn
	if command.TypeSelector(buffer) != command.CARD_REQUEST {
		return
	}
	if logic.ValidTest(command.CardsDeSerialize(buffer[2:])) != logic.INVALID {
		sendCardResponseOK(fisrt)
	} else {
		sendCardResponseNOK(fisrt)
		putCardsCommand(fisrt)
		goto start
	}

	destop = command.CardsDeSerialize(buffer[2:])
	tmp = other
	tmp.ChannelOut <- command.DestopCardsSeriable(destop)
	for {
		putCardsCommand(tmp)

		buffer = <-tmp.ChannelIn
		log.Println(buffer)
		if command.TypeSelector(buffer) == command.PASS_REQUSET {
			log.Println("passing...")
			//sendCardResponseOK(tmp)
			if tmp == fisrt {
				tmp = other
			} else {
				tmp = fisrt
			}
			destop = nil
			continue
		}

		if command.TypeSelector(buffer) != command.CARD_REQUEST {
			return
		}
		newCards = command.CardsDeSerialize(buffer[2:])

		if logic.Compare(destop, newCards) == logic.SECOND {
			sendCardResponseOK(tmp)
			log.Println("card is valid")
			if tmp == fisrt {
				tmp = other

			} else {
				tmp = fisrt
			}
			destop = newCards
			tmp.ChannelOut <- command.DestopCardsSeriable(destop)
		} else {
			log.Println("card is invalid")
			sendCardResponseNOK(tmp)
		}

	}

}

func disideCardFirst(c1, c2 *client.Client) (*client.Client, *client.Client) {
	i := rand.Intn(2)
	if i == 0 {
		c1.ChannelOut <- command.PutCardSerialize()
		return c1, c2
	} else {
		c2.ChannelOut <- command.PutCardSerialize()
		return c2, c1
	}
}

func dealCards(cs []cards.Card, c *client.Client) {
	c.ChannelOut <- command.DealCardsSerialize(cs)
}

func putCardsCommand(c *client.Client) {
	c.ChannelOut <- command.PutCardSerialize()
}

/* 出牌和过牌都用command.Response来响应 */
func sendCardResponseOK(c *client.Client) {
	c.ChannelOut <- command.ResponseSerialize(command.RESULT_OK, command.CARD_RESPONSE)
}

func sendCardResponseNOK(c *client.Client) {
	c.ChannelOut <- command.ResponseSerialize(command.RESULT_NOK, command.CARD_RESPONSE)
}

func sendPassResponseOK(c *client.Client) {
	c.ChannelOut <- command.ResponseSerialize(command.RESULT_OK, command.PASS_RESPONSE)
}

func sendPassResponseNOK(c *client.Client) {
	c.ChannelOut <- command.ResponseSerialize(command.RESULT_NOK, command.PASS_RESPONSE)
}

func Bussiness2(c1, c2 *client.Client) {
	cards1, cards2 := cards.GenerateCard()

	var cards_in_desktop, cards_new []cards.Card
	var buffer string
	var tmp *client.Client
	/* 给客户端建立连接的时间 */
	time.Sleep(time.Second * 2)
	dealCards(cards1, c1)
	dealCards(cards2, c2)
	fisrt, second := disideCardFirst(c1, c2)

start:
	buffer = <-fisrt.ChannelIn
	if command.TypeSelector(buffer) != command.CARD_REQUEST {
		return
	}
	if logic.ValidTest(command.CardsDeSerialize(buffer[2:])) != logic.INVALID {
		sendCardResponseOK(fisrt)
	} else {
		sendCardResponseNOK(fisrt)
		putCardsCommand(fisrt)
		goto start
	}

	getOtherClient := func(tmp *client.Client) *client.Client {
		if tmp == fisrt {
			return second
		} else {
			return fisrt
		}
	}

	exchangeClient := func(tmp_pointer **client.Client) {
		if *tmp_pointer == fisrt {
			*tmp_pointer = second
		} else {
			*tmp_pointer = fisrt
		}
	}

	var (
		idle          uint8 = 0
		get_cards     uint8 = 1
		cards_invalid uint8 = 2
		cards_valid   uint8 = 3
	)
	var state uint8 = idle
	var illegal_count int = 0

	cards_in_desktop = command.CardsDeSerialize(buffer[2:])
	tmp = second
	tmp.ChannelOut <- command.DestopCardsSeriable(cards_in_desktop)
	for {
		switch state {
		case idle:
			putCardsCommand(tmp)
			state = get_cards
		case get_cards:
			buffer = <-tmp.ChannelIn
			if command.TypeSelector(buffer) == command.CARD_REQUEST {
				cards_new = command.CardsDeSerialize(buffer[2:])
				if logic.Compare(cards_in_desktop, cards_new) == logic.SECOND {
					sendCardResponseOK(tmp)
					cards_in_desktop = cards_new
					log.Println("cards are valid")
					state = cards_valid
				} else {
					sendCardResponseNOK(tmp)
					state = cards_invalid
					log.Println("card are invalid")
				}
			} else if command.TypeSelector(buffer) == command.PASS_REQUSET {
				/* 防止反复pass */
				if cards_in_desktop == nil {
					sendPassResponseNOK(tmp)
				} else {
					sendPassResponseOK(tmp)
					exchangeClient(&tmp)
					/* 清空牌桌 */
					cards_in_desktop = nil
				}
				state = idle
			} else {
				log.Println("illegal request")
				illegal_count++
				if illegal_count > 10 {
					return
				}
				state = idle
			}

		case cards_invalid:
			state = idle
		case cards_valid:
			other := getOtherClient(tmp)
			other.ChannelOut <- command.DestopCardsSeriable(cards_in_desktop)
			exchangeClient(&tmp)
			state = idle
		}
	}

}
