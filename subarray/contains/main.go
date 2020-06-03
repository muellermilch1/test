package main

import "fmt"

//IsValidSubsequence returns if seq is a part of array
func IsValidSubsequence(array []int, sequence []int) bool {
	ssc := 0
	for _, v := range array {
		if v == sequence[ssc] {
			ssc++
		}
		if ssc == len(sequence) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(IsValidSubsequence([]int{1, 1, 6, 1}, []int{1, 1, 1, 6}))
}
