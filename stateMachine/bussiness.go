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
		sendOK(fisrt)
	} else {
		sendNOK(fisrt)
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
			//sendOK(tmp)
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
			sendOK(tmp)
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
