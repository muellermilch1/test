package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	fmt.Println(haystack)
	for i := 0; i < len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			break
		}
		for j := range needle {
			fmt.Println(j, j+i)
			if needle[j] == haystack[i+j] && j == len(needle)-1 {
				return i
			} else if needle[j] != haystack[i+j] {
				break
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hel", "ll"))
}
