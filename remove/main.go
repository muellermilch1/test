package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func removeKdigits(num string, k int) string {
	if len(num) <= k {
		return "0"
	}
	if k == 0 {
		return num
	}
	rest := getVariations(num)
	in := smallest(rest, k)
	fmt.Println(in)
	var ret string
	for i := range num {
		if _, ok := in[i]; !ok && !(len(ret) == 0 && num[i] == '0') {
			ret = ret + string(num[i])
		}
	}
	return ret
}

func smallest(rest []int, k int) map[int]int {
	sorted := make([]int, len(rest))
	copy(sorted, rest)
	sort.Ints(sorted)
	ret := make(map[int]int)
	for i := 0; i < k; i++ {
		for j, v := range rest {
			if v == sorted[i] {
				ret[j] = v
				rest[j] = -1
				break
			}
		}
	}
	return ret
}

func getVariations(num string) []int {
	var rest []int
	for i := range num {
		r := num[:i] + num[i+1:]
		n, err := strconv.Atoi(r)
		if err != nil {
			log.Fatal(err)
		}
		rest = append(rest, n)
	}
	return rest
}

func main() {
	fmt.Println(removeKdigits("1234567890", 9))
}
