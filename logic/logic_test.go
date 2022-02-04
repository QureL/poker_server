package logic_test

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	var marks int = -1

	switch {
	case marks == 1:
		fmt.Println("1")
	case marks > 1:
		fmt.Print(">1")
	}
	fmt.Println("j")

}
