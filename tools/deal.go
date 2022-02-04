package tools

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

/*

func GenerateDiscard() []int {
	var result []int

	for len(result) < cards.DISCARD_NUM {
		i := rand.Intn(cards.DECK_SIZE)
		flag := true
		for _, tmp := range result {
			if tmp == i {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, i+1)
		}
	}
	return result
}
*/

/** Returns a random shuffling of the array. */
func Shuffle(nums []int) []int {
	len := len(nums)
	cur := make([]int, len)
	for i := 0; i < len; i++ {
		cur[i] = nums[i]
	}
	var (
		pos  int
		temp int
	)
	for i := len - 1; i >= 0; i-- {
		pos = rand.Intn(i + 1)
		temp = cur[pos]
		cur[pos] = cur[i]
		cur[i] = temp
	}
	return cur
}
