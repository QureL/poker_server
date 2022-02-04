package tools_test

import (
	"fmt"
	"poker_server/cards"
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
}
