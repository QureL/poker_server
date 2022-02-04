package statemachine

import (
	"math/rand"
	"poker_server/cards"
	"poker_server/client"
	"poker_server/command"
	"poker_server/logic"
)

func Bussiness(c1, c2 *client.Client) {
	cards1, cards2 := cards.GenerateCard()

	var destop, newCards []cards.Card
	var buffer string
	var tmp *client.Client

	dealCards(cards1, c1)
	dealCards(cards2, c2)
	fisrt, other := disideCardFirst(c1, c2)

	buffer = <-fisrt.ChannelIn
	if command.TypeSelector(buffer) != command.CARD_REQUEST {
		return
	}

	destop = command.CardsDeSerialize(buffer[2:])
	tmp = other
	for {
		putCardsCommand(tmp)

		buffer = <-tmp.ChannelIn
		if command.TypeSelector(buffer) != command.CARD_REQUEST {
			return
		}
		newCards = command.CardsDeSerialize(buffer[2:])

		if logic.Compare(destop, newCards) != logic.SECOND {
			sendOK(tmp)
			if tmp == fisrt {
				tmp = other
			} else {
				tmp = fisrt
			}
			destop = newCards
		} else {
			sendNOK(tmp)
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

func sendOK(c *client.Client) {
	c.ChannelOut <- command.ResponseSerialize(command.RESULT_OK, command.CARD_RESPONSE)
}

func sendNOK(c *client.Client) {
	c.ChannelOut <- command.ResponseSerialize(command.RESULT_NOK, command.CARD_RESPONSE)
}
