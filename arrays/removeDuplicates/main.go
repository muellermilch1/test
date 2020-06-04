package main

import "fmt"

// func removeDuplicates(nums []int) int {
// 	if len(nums) <= 1 {
// 		return len(nums)
// 	}
// 	last := nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		if last == nums[i] {
// 			nums = append(nums[:i], nums[i+1:]...)
// 			i--
// 		} else {
// 			last = nums[i]
// 		}
// 	}
// 	return len(nums)
// }

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	last := nums[0]
	high, low := 1, 1
	for high < len(nums) {
		if nums[high] != last {
			nums[low] = nums[high]
			last = nums[high]
			low++
		}
		high++
	}
	return low
}

func main() {
	a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(a))
	fmt.Println(a)
}
