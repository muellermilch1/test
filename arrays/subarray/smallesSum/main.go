package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(s int, nums []int) int {
	low, high := 0, 0
	sa := math.MaxInt64
	for high < len(nums) {
		a := sum(nums, low, high)
		if a == s && high-low+1 < sa {
			sa = high - low + 1
			low++
		} else if a > s && high-low > 0 {
			low++
		} else {
			high++
		}
	}
	if sa == math.MaxInt64 {
		return 0
	}
	return sa
}

func sum(nums []int, low, high int) int {
	var sum int
	for i := low; i <= high; i++ {
		sum += nums[i]
	}
	return sum
}

func main() {
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
}
