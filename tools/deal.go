package tools

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

/** Knuth-Durstenfeld */
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
