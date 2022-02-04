package cards_test

import (
	"fmt"
	"poker_server/cards"
	"testing"
)

func TestGenerate(t *testing.T) {
	arr1, arr2 := cards.GenerateCard()

	fmt.Println(len(arr1), arr1)
	fmt.Println(len(arr2), arr2)
}
