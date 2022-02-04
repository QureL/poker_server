package statemachine_test

import (
	"fmt"
	"testing"
	"time"
)

func oTestChannel(t *testing.T) {
	c := make(chan struct{})

	go func() {
		fmt.Println("goroutine start...")

		for {
			select {
			case <-c:
				fmt.Println("case1")
				goto end
			default:
				fmt.Println("default..")
			}
		}
	end:
		fmt.Println("goroutine end ...")
	}()

	time.Sleep(time.Second * 2)

	c <- struct{}{}
	time.Sleep(time.Second)
	fmt.Println("main end...")
}

func oTestChannel02(t *testing.T) {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	//c2 <- "hello"

	select {
	case msg1 := <-c1:
		fmt.Println("c1 received: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 received: ", msg2)
	default:
		//	fmt.Println("No data received.")
	}
}

func pTestChannel03(t *testing.T) {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	// 开启一个协程，可以发送数据到信道
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "hello"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("c1 received: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 received: ", msg2)
	}
	fmt.Println("main end....")
}

func TestChannel04(t *testing.T) {
	c := make(chan struct{})

	go func() {
		for {
			flag := false
			select {
			case <-c:
				fmt.Println("case shoot")
				flag = true
			default:
				fmt.Println("default...")
			}
			if flag {
				fmt.Println("flag true")
				break
			}
		}
		fmt.Println("goroutine end...")
	}()

	time.Sleep(time.Second)
	c <- struct{}{}

	time.Sleep(time.Second)

	fmt.Println("main end...")
}
