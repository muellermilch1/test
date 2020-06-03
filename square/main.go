package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime/trace"
)

func isPerfectSquare(num int) bool {
	f, err := os.Create("out.trace")
	if err != nil {
		log.Fatal(err)
	}
	trace.Start(f)
	defer trace.Stop()
	s := int(math.Sqrt(float64(num)))
	return num == s*s
}

func main() {
	fmt.Println(isPerfectSquare(16))
}
