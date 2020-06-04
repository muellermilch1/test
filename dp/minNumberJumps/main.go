package main

import (
	"fmt"
	"math"
)

//MinNumberOfJumps returns the minimal number of jumps
//needed to get to the end of the array
func MinNumberOfJumps(array []int) int {
	dp := make([]int, len(array))
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}
	for i, v := range array {
		for j := 1; j <= v && j+i < len(array); j++ {
			dp[i+j] = min(dp[i+j], dp[i]+1)
		}
	}
	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//for demonstration
func main() {
	fmt.Println(MinNumberOfJumps([]int{3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3}))
}
